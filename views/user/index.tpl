<!DOCTYPE html>

<html>
    <head>
        <title>Beego</title>
        <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
        <!-- 最新版本的 Bootstrap 核心 CSS 文件 -->
        <link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/3.4.1/css/bootstrap.min.css" integrity="sha384-HSMxcRTRxnN+Bdg0JdbxYKrThecOKuH5zCYotlSAcp1+c8xmyTe9GYg1l9a69psu" crossorigin="anonymous">
    </head>
    <body>
        <header class="hero-unit" style="background-color:#A9F16C">
            <div class="container">
                <div class="row">
                    <div class="hero-text">
                        <h1>System Infomation</h1>
                        <p class="description">
                            Download System Infomation
                        </p>
                    </div>
                </div>
                <nav class="navbar navbar-default" role="navigation">
                    <div class="container-fluid">
                        <div class="navbar-header">
                            <a class="navbar-brand" href="/">Report</a>
                        </div>
                        <ul class="nav navbar-nav navbar-right">
                            <li><a href="/"><span class="glyphicon glyphicon-file"></span> Report</a></li>
                            <li><a href="/device"><span class="glyphicon glyphicon-folder-open"></span> Device</a></li>
                            <li><a href="/userinfo"><span class="glyphicon glyphicon-wrench"></span> userinfo</a></li>
                        </ul>
                    </div>
                </nav>
            </div>
        </header>
        <div class="panel panel-default">
            <div class="panel-body">
                <ul class="list-group">
                {{range $index,$elem := .Files}}
                    <li class="list-group-item">
                        <a download="" href="/file/down/{{$elem}}">
                            {{$elem}}
                        </a>
                    </li>
                {{end}}
                </ul>
            </div>
        </div>
    </body>
</html>