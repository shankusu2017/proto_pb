# proto_pb
protocol define use protobuf

# pb 为 protobuf 格式的协议文件
# 预备了其它格式的协议文件的可能

# 生成 go 文件的参考命令（pwd==当前目录）
# 在pb目录下执行
protoc --go_out=../go/proto *.proto