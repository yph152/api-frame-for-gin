#### goimports 是 Go 语言官方提供的工具，它能够为我们自动格式化 Go 语言代码并对所有引入的包进行管理，包括自动增删依赖包的引用、将依赖包按字母排序并分类；goimports = gofmt + 依赖包管理；
#### golint + golangci-lint:

## golang 目录规范
```
├── api
├── assets
├── build
├── cmd
├── configs
├── deployments
├── docs
├── examples
├── githooks
├── init
├── internal
├── LICENSE.md
├── Makefile
├── pkg
├── README.md
├── scripts
├── test
├── third_party
├── tools
├── vendor
├── web
└── website
```

### Golang 的目录：

#### ```/cmd```:

该项目的主要目录。

每个应用程序的目录名称应与您想要的可执行文件的名称相匹配 （例如：```/cmd/myapp```）

不要在应用程序目录中放入大量的代码。如果您认为代码可以导入并在其它项目中使用，那么它应该存在于```/pkg```目录中。如果代码不可重用或者您不希望其他人重用它，请将该代码放在```/internal``` 目录中。你会惊讶于别人会做什么，所以要明确你的意图！

通常有一个小```main```函数可以从```internal```和```pkg```目录中导出和调用代码，而不是其它任何东西。

#### ```/internal```

私有应用程序和库代码。这是您不希望其他人在应用程序或库中导入的代码。

将您的实际应用程序代码放在```/internal/app```目录( 例如```/internal/app/myapp```) 和 ```/internal/pkg```目录中这些应用程序共享的代码（例如： /internal/pkg/myprivlib）。

#### ```/pkg```

可以由外部应用程序使用的库代码（/pkg/mypluliclib）。其它项目将导入这些库，期望它们可以工作，所以在你把东西放在这里之前要三思而后行。

当你的根目录包含许多非Go组件和目录时，它也可以在一个地方将 Go 代码分组，从而更容易运行各种Go工具；

#### ```/vendor```

应用程序依赖项（手动管理或由您喜欢的依赖管理工具dep）

### 服务应用程序目录

#### ```/api```
OpenAPI / Swagger规范， JSON模式文件，协议定义文件。

#### ```/Web```
特定于Web应用程序的组件： 静态Web资产，服务器端模板和SPA。

### 常见的应用程序目录

#### ```/configs```
配置文件模板或默认配置。
将您的```confd```或者```consul-template```模板文件放在这里。

#### ```/init```
系统初始化（systemd, upstart）和进程管理器/(runit,supervisord) 配置。

#### ```/scripts```
脚本执行各种构建，安装，分析等操作。

这些脚本使根级 Makefile 保持简洁。

#### ```/build```
包装和持续集成。
将您的云(AMI),容器（Docker）,OS(dep,rpm,pkg)包配置和脚本放在 ```/build/package```目录中。

将CI(travis,circle,drone) 配置和脚本放在```/build/ci```目录中。


####  	```/deployments```
IaaS, Pass, 系统和容器编排部署配置和模板（docker-compose, kubernetes/helm）。

#### ```/test```
其它外部测试应用和测试数据。您可以随意构建```/test```目录。对于更大的项目，有一个数据子目录是有意义的。例如，您可以拥有```/test/data```或者```/test/testdata```如果需要Go来忽略该目录中的内容。请注意，Go也会忽略以". "开头的目录或文件。或“_”，因此您在命名测试数据目录方面具有更大的灵活性。

### 其它目录
#### ```/docs```
设计和用户文档（除了你的 godoc 生成的文档）。

#### ```/tools```
该项目的支持工具。请注意，这些工具可以从 /pkg 和 /internal 目录中导入代码。

#### ```/examples```
应用程序或公共库的示例。

#### ```/third_party```
外部帮助工具, 分叉代码和其它第三方使用程序（例如，Swagger UI)。

####  ```/githooks```
Git 钩子。

#### ```/assets```
与您的存储克一起使用的其它资产(图像，徽标等)。

#### ```/website```
如果您不使用Github页面，这是放置项目的网站数据的地方。

### 你不应该有的目录
#### ```/src```
一些Go项目确实有一个```/src```文件夹，但它通常发生在开发人员来自Java世界时，它是一种常见的模式。如果您可以帮助自己尝试不采用此Java模式。你真的不希望你的Go代码或Go项目看起来像Java。



 
### golang 代码规范
### golang api 规范