# go语言简洁架构

## 目录

### 参考献文

- 标准Go项目布局:https://github.com/golang-standards/project-layout
- go语言ddd实践:https://github.com/marcusolsson/goddd
- Go中的清洁架构:https://github.com/eminetto/clean-architecture-go
- 上个仓库的演化版本:https://github.com/eminetto/clean-architecture-go



### `/cmd`

该项目的主要应用。

每个应用程序的目录名称应与您想要的可执行文件的名称相匹配（例如/cmd/myapp）。

不要在应用程序目录中放入大量代码。如果您认为代码可以导入并在其他项目中使用，那么它应该存在于`/pkg`目录中。
如果代码不可重用或者您不希望其他人重用它，请将该代码放在`/internal`目录中。你会惊讶于别人会做什么，所以要明确你的意图！

通常有一个小`main`函数可以从`/internal`和`/pkg`目录中导入和调用代码，而不是其他任何东西。

请参阅`/cmd`目录以获取示例。

### `/internal`

私有应用程序和库代码。这是您不希望其他人在其应用程序或库中导入的代码。

将您的实际应用程序代码放在`/internal/app`目录（例如`/internal/app/myapp`）和`/internal/pkg`目录中这些应用程序共享的代码（例如`/internal/pkg/myprivlib`）。

### `/pkg`

可以由外部应用程序使用的库代码（例如`/pkg/mypubliclib`）。其他项目将导入这些库，期望它们可以工作，所以在你把东西放在这里之前要三思而后行:-)

当你的根目录包含许多非Go组件和目录时，它也可以在一个地方将Go代码分组，从而更容易运行各种Go工具（如Best Practices for Industrial ProgrammingGopherCon EU 2018中所述）。

`/pkg`如果您想查看哪个热门的Go repos使用此项目布局模式，请查看该目录。这是一种常见的布局模式，但它并未被普遍接受，Go社区中的一些人不推荐它。

### `/vendor`

应用程序依赖项（手动管理或由您喜欢的依赖管理工具管理dep）。

如果要构建库，请不要提交应用程序依赖项。

## 服务应用程序目录

### `/api`
`OpenAPI/Swagger`规范，`JSON`模式文件，协议定义文件。

请参阅`/api`目录以获取示例。

## Web应用程序目录
### `/web`
特定于Web应用程序的组件：静态Web资产，服务器端模板和SPA。

## 常见应用程序目录
### `/configs`
配置文件模板或默认配置。

将您的`confd`或`consul-template`模板文件放在这里。

### `/init`
系统初始化（systemd，upstart，sysv）和进程管理器/主管（runit，supervisord）配置。

### `/scripts`
脚本执行各种构建，安装，分析等操作。

这些脚本使根级Makefile保持简洁（例如https://github.com/hashicorp/terraform/blob/master/Makefile）。

请参阅/scripts目录以获取示例。

### `/build`
包装和持续集成。

将您的云（AMI），容器（Docker），OS（deb，rpm，pkg）包配置和脚本放在/build/package目录中。

将CI（travis，circle，drone）配置和脚本放在/build/ci目录中。请注意，某些CI工具（例如，Travis CI）对其配置文件的位置非常挑剔。尝试将配置文件放在/build/ci将它们链接到CI工具所期望的位置的目录中（如果可能）。

### `/deployments`
IaaS，PaaS，系统和容器编排部署配置和模板（docker-compose，kubernetes / helm，mesos，terraform，bosh）。

### `/test`
其他外部测试应用和测试数据。您可以随意构建/test目录。对于更大的项目，有一个数据子目录是有意义的。例如，您可以拥有/test/data或者/test/testdata如果需要Go来忽略该目录中的内容。请注意，Go也会忽略以“。”开头的目录或文件。或“_”，因此您在命名测试数据目录方面具有更大的灵活性。

请参阅`/test`目录以获取示例。

## 其他目录
### `/docs`
设计和用户文档（除了你的godoc生成的文档）。

请参阅/docs目录以获取示例。

### `/tools`
该项目的支持工具。请注意，这些工具可以从/pkg和/internal目录中导入代码。

请参阅/tools目录以获取示例。

### `/examples`
应用程序和/或公共库的示例。

请参阅`/examples`目录以获取示例。

### `/third_party`
外部帮助工具，分叉代码和其他第三方实用程序（例如，Swagger UI）。

### `/githooks`
Git钩子。

### `/assets`
与您的存储库一起使用的其他资产（图像，徽标等）。

### `/website`
如果您不使用Github页面，这是放置项目的网站数据的地方。



