# indexer

存储向量库时候使用，把文档和文档的向量存储到后端存储系统，需要向量型数据库，不是普通的mysql那种关系型数据库

demo用Milvus原生支持，用docker启动

```
0. 初始和创建emberrer
1. 创建索引器
2. 索英文当
```

可以看Milvus的官方文档https://milvus.io/docs/zh/install-overview.md，但是感觉中文是机翻的，要写一个go的客户端去链接然后存储数据

> ### 通过 PowerShell 或 Windows 命令提示符                        
>
> 如果您更熟悉 PowerShell 或 Windows 命令提示符，请按以下步骤操作。
>
> 1. 右键单击 Docker Desktop 并选择**“以管理员身份运行**”，以管理员模式打开 Docker Desktop。
>
> 2. 下载安装脚本，并将其保存为`standalone.bat` 。
>
>    ```powershell
>    C:\>Invoke-WebRequest https://raw.githubusercontent.com/milvus-io/milvus/refs/heads/master/scripts/standalone_embed.bat -OutFile standalone.bat
>    ```
>
> 3. 运行下载的脚本，以 Docker 容器的形式启动 Milvus。
>
>    ```powershell
>    C:\>standalone.bat start
>    Wait for Milvus starting...
>    Start successfully.
>    To change the default Milvus configuration, edit user.yaml and restart the service.
>    ```
>
>    运行安装脚本后：
>
>    - 一个名为**Milvus Standalone 的**Docker 容器已在端口**19530** 上启动。
>    - 一个嵌入式 etcd 已随 Milvus 一起安装在同一容器中，并在端口**2379** 上提供服务。其配置文件映射到当前文件夹中的**embedEtcd.yaml**。
>    - Milvus 数据卷已映射到当前目录下的**volumes/milvus**目录中。
>
>    您可以使用以下命令管理 Milvus 容器及存储的数据。
>
>    ```powershell
>    # Stop Milvus
>    C:\>standalone.bat stop
>    Stop successfully.
>    
>    # Delete Milvus container
>    C:\>standalone.bat delete
>    Delete Milvus container successfully. # Container has been removed.
>    Delete successfully. # Data has been removed.
>    ```

脚本已经下好在改目录里了，直接启动，启动报错就在前面加`./`前缀

诶诶可恶的powershell（我要重装电脑）

访问`http://127.0.0.1:9091/webui/` 会有一个web页面，相当于一个管理端，但是看不到更具体的一些数据，看到官方文档发现Attu可以作为管理milvus的webAPP

看到attu的release界面也是支持docker启动的

```
docker run -d --name attu -p 3000:3000 -v attu-data:/data zilliz/attu:v3.0.0-beta.6
```

肥肠方便啊

然后在写代码的时候发现官方文档的代码例子是错的。。。

这里注意attu连接milvus的时候**不要填** `localhost:19530` 或 `127.0.0.1:19530`，因为 Attu 前端在浏览器里跑，浏览器会把 `localhost` 解析成它自己，而不是你的电脑