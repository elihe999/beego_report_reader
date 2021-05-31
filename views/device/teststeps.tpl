<!DOCTYPE html>

<html>
    <head>
        <title>Beego</title>
        <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
        <link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/3.4.1/css/bootstrap.min.css" integrity="sha384-HSMxcRTRxnN+Bdg0JdbxYKrThecOKuH5zCYotlSAcp1+c8xmyTe9GYg1l9a69psu" crossorigin="anonymous">
    </head>
    <body>
        <header class="hero-unit" style="background-color:#A9F16C">
            <div class="container">
                <div class="row">
                    <div class="hero-text">
                        <h1>Device Test Step</h1>
                        <p class="description">
                            Select a device to see the test details.
                        </p>
                    </div>
                </div>
                <nav class="navbar navbar-default" role="navigation"> 
                    <div class="container-fluid">
                        <div class="navbar-header"> 
                            <a class="navbar-brand" href="/">Report</a> 
                        </div>
                        <ul class="nav navbar-nav navbar-right"> 
                            <li><a href="/sip"><span class="glyphicon glyphicon-file"></span> SIP</a></li> 
                            <li><a href="/device"><span class="glyphicon  glyphicon-folder-open"></span> Device</a></li> 
                            <li><a href="/userinfo"><span class="glyphicon glyphicon-wrench"></span> userinfo</a></li> 
                        </ul>
                    </div>
                </nav>
            </div>
        </header>
        <div class="panel panel-default">
            <div class="panel-body">
                {{range $elem := .Files}}
                    <p>
                        {{$elem}}
                    </p>
                {{end}}
            </div>
        </div>
    </body>
</html>