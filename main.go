// Copyright 2021 wangkun.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package main

import (
	"github.com/karry-11110/LM_system/app/console"
	"github.com/karry-11110/LM_system/app/http"
	"github.com/karry-11110/ginx/framework"
	"github.com/karry-11110/ginx/framework/provider/app"
	"github.com/karry-11110/ginx/framework/provider/cache"
	"github.com/karry-11110/ginx/framework/provider/config"
	"github.com/karry-11110/ginx/framework/provider/distributed"
	"github.com/karry-11110/ginx/framework/provider/env"
	"github.com/karry-11110/ginx/framework/provider/id"
	"github.com/karry-11110/ginx/framework/provider/kernel"
	"github.com/karry-11110/ginx/framework/provider/log"
	"github.com/karry-11110/ginx/framework/provider/orm"
	"github.com/karry-11110/ginx/framework/provider/redis"
	"github.com/karry-11110/ginx/framework/provider/ssh"
	"github.com/karry-11110/ginx/framework/provider/trace"
)

func main() {
	// 初始化服务容器
	container := framework.NewGinxContainer()
	// 绑定App服务提供者
	container.Bind(&app.GinxAppProvider{})
	// 后续初始化需要绑定的服务提供者...
	container.Bind(&env.GinxEnvProvider{})
	container.Bind(&distributed.LocalDistributedProvider{})
	container.Bind(&config.GinxConfigProvider{})
	container.Bind(&id.GinxIDProvider{})
	container.Bind(&trace.GinxTraceProvider{})
	container.Bind(&log.GinxLogServiceProvider{})
	container.Bind(&orm.GormProvider{})
	container.Bind(&redis.RedisProvider{})
	container.Bind(&cache.GinxCacheProvider{})
	container.Bind(&ssh.SSHProvider{})

	// 将HTTP引擎初始化,并且作为服务提供者绑定到服务容器中
	if engine, err := http.NewHttpEngine(container); err == nil {
		container.Bind(&kernel.GinxKernelProvider{HttpEngine: engine})
	}

	// 运行root命令
	console.RunCommand(container)
}
