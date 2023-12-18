### Inscription tool
#### 介绍
高性能evm系铭刻工具，用于铭刻交易的发送，支持并发发送
#### 安装
```shell
cd inscriptiontool && go install .
```

#### 使用
* 创建钱包
```shell
inscriptiontool wallet create -n <wallet name>
```
* 导入钱包
```shell
inscriptiontool wallet restore -n <wallet name> -p <private key>
```
* 获取钱包列表
```shell
inscriptiontool wallet list
```
* 删除钱包
```shell
inscriptiontool wallet delete -n <wallet name>
```

#### 发送铭刻交易
```shell
# 查看帮助
inscriptiontool run --help
Usage:
  inscriptiontool run [flags]

Flags:
  -a, --account string   账户名称
  -m, --amount float     转账金额, (默认是0)
  -c, --count int        重复次数 (default 1)
  -d, --data string      data数据
  -g, --gf float         gas倍数 (default 1)
  -l, --glf int          gasLimit倍数, 默认是2，即21000*2 (default 2)
  -h, --help             help for run
  -r, --rpcUrl string    rpc 地址
  -t, --to string        转账To地址, (默认是转给自己)

# 开启铭刻任务
inscriptiontool run -a <wallet name> -m <amount> -t <to address> -r <url> -c <count> -d <data> -g <gas factor> -l <gas limit factor>
# 使用示例
inscriptiontool run -a whf -d 'data:,{"p":"bsc-20","op":"mint","tick":"sofi","amt":"4"}' -r 'https://bsc.blockpi.network/v1/rpc/' -c 100
```
