# go-ratel

[Ratel](https://github.com/ainilili/ratel) 命令行斗地主 `golang`客户端，完全支持Netty + Protobuf的通信协议，100%兼容`Ratel`服务端。

## 关键点剖析

**1.使用Google Protobuf - 实现跨平台跨语言的序列化/反序列化**

**2.Ratel使用Netty内置的Protobuf编解码器**

- 编码器：ProtobufVarint32LengthFieldPrepender

```
BEFORE DECODE (302 bytes)       AFTER DECODE (300 bytes)
+--------+---------------+      +---------------+
| Length | Protobuf Data |----->| Protobuf Data |
| 0xAC02 |  (300 bytes)  |      |  (300 bytes)  |
+--------+---------------+      +---------------+
```

- 解码器：ProtobufVarint32FrameDecoder

```
BEFORE ENCODE (300 bytes)       AFTER ENCODE (302 bytes)
+---------------+               +--------+---------------+
| Protobuf Data |-------------->| Length | Protobuf Data |
|  (300 bytes)  |               | 0xAC02 |  (300 bytes)  |
+---------------+               +--------+---------------+
```

大致的协议就是，一个完整包由`头部（Header）`和`数据体（Body）`组成。`编码器`会计算出`Body`的长度放在`Header`中。`Body`的长度
采用 [Varint](https://developers.google.com/protocol-buffers/docs/encoding#varints)，会把整数编码为变长字节。对于
32位整型数据经过Varint编码后需要1~5个字节，小的数字使用1个byte，大的数字使用5个bytes。

**3.Golang客户端处理粘包和拆包**

- 编码：序列化结构体，计算数据包长度，换算成Varint

```go
func (c *Codec) Encode(transferData *common.ServerTransferDataProtoc, duration duration.Duration) error {
	// protobuf 序列化
	encodeData, e := proto.Marshal(transferData)
	if e != nil {
		return e
	}
	// 计算数据体长度
	bodyLen := len(encodeData)
	if bodyLen > MaxContextLen {
		return errors.New("not enough")
	}
	// 使用Varint类型
	header := proto.EncodeVarint(uint64(bodyLen))

	buffer := make([]byte, len(header) + bodyLen)
	copy(buffer, header)
	copy(buffer[len(header):], encodeData)

	_, e = c.Conn.Write(buffer)
	return e
}
```

- 解码：计算数据包头部，换算Varint，处理粘包，反序列结构体

```go
func (c *Codec) Decode() (*common.ServerTransferDataProtoc, bool, error) {
	// 计算出 body 体长度，以及头部占用的字节数 size
	bodyLen, size := proto.DecodeVarint(c.Buffer.buf[c.Buffer.start:])
	if bodyLen > MaxContextLen {
		return nil, false, errors.New("not enough")
	}
	if bodyLen == 0 {
		return nil, false, nil
	}
	// 在当读取的字节数不满足 body 体的长度，进入下一轮
	body, e := c.Buffer.read(size, int(bodyLen))
	if e != nil {
		return nil, false, nil
	}

	transferData := common.ServerTransferDataProtoc{}
	// 反序列化
	e = proto.Unmarshal(body, &transferData)
	if e != nil {
		return nil, false, e
	}
	return &transferData, true, nil
}
```

## 参考资料

1.[详解varint编码原理](https://segmentfault.com/a/1190000020500985?utm_source=tag-newest)

2.[Netty源码分析-ProtobufVarint32FrameDecoder](https://blog.csdn.net/nimasike/article/details/101392803)

3.[golang 使用 protobuf 的教程](https://www.cnblogs.com/smallleiit/p/10926794.html)