<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/3.4.1/css/bootstrap.min.css" integrity="sha384-HSMxcRTRxnN+Bdg0JdbxYKrThecOKuH5zCYotlSAcp1+c8xmyTe9GYg1l9a69psu" crossorigin="anonymous">
    <title>Mscgen diagram</title>
    <script src="/static/mscgen-inpage.js" defer></script>
    <script type="text/javascript" src="/static/axios.min.js"></script>
    <style>
        body,html {
            height: 100%;
        }
        #showarea {
            width: 600px;
            height: 830px;
        }
        #box {
            height: 850px;
            width: 100%;
            overflow: hidden;
        }
        #resize {
            width: 5px;
            height: 850px;
            cursor: w-resize;
            float: left;
            background: skyblue;
        }
        #left {
            width: calc(50% - 5px);
            height: 100%;
            float: left;
        }
        #right {
            float: left;
            height: 100%;
        }
        textarea {
            font-family: Consolas,Monaco,monospace;
        }
    </style>
</head>
<body>
    <header class="hero-unit" style="background-color:#A9F16C">
        <div class="container">
            <div class="row">
                <div class="hero-text">
                    <h1>Report</h1>
                    <p class="description">
                        Test Case Result
                    </p>
                </div>
            </div>
            <nav class="navbar navbar-default" role="navigation"> 
                <div class="container-fluid">
                    <div class="navbar-header"> 
                        <a class="navbar-brand" href="#">Report</a> 
                    </div> 
                    <ul class="nav navbar-nav navbar-right"> 
                        <li><a href="sip"><span class="glyphicon glyphicon-file"></span> SIP</a></li> 
                        <li><a href="device"><span class="glyphicon  glyphicon-folder-open"></span> Device</a></li> 
                        <li><a href="userinfo"><span class="glyphicon glyphicon-wrench"></span> userinfo</a></li> 
                    </ul>
                </div> 
            </nav>
        </div>
    </header>
    <div id="box">
        <div id="left">
            <pre id="mscremote" class="mscgen_js" style="max-height: 850px;">
                msc {
                    wordwraparcs="true";
                    {{range $elem := .mscbody}}
                        {{$elem}}
                    {{end}}
                }
            </pre>
        </div>
        <div id="resize"></div>
        <div id="right">
            <textarea id="showarea" readonly>
            </textarea>
        </div>
    </div>
</body>
<script>
    var path = "11"
    window.onload = function () {
        var resize = document.getElementById("resize")
        var left = document.getElementById("left")
        var right = document.getElementById("right")
        var box = document.getElementById("box")
        var showbox = document.getElementById("showarea")
        resize.onmousedown = function (e) {
            var startX = e.clientX
            resize.left = resize.offsetLeft
            document.onmousemove = function (e) {
                var endX = e.clientX;
                var moveLen = resize.left + (endX - startX)
                var maxT = box.clientWidth - resize.offsetWidth
                if (moveLen < 400) {
                    moveLen = 400
                } else if (moveLen > (maxT-340)) {
                    moveLen = maxT-340
                }
                console.log(moveLen)
                resize.style.left = moveLen + 1
                left.style.width = moveLen + "px"
                right.style.width = (box.clientWidth - moveLen - 5) + "px"
                showbox.style.width = (box.clientWidth - moveLen - 5) + "px"
                showbox.style.resuze = "horizontal"
            }
            document.onmouseup = function (evt) {
                document.onmousemove = null
                document.onmouseup = null
                resize.releaseCapture && resize.releaseCapture()
            }
            resize.setCapture && resize.setCapture()
            return false
        }
    }
    //todo
    function show(a) {
        var pDiv = document.getElementById("showarea")
        console.log("test")
    }
</script>

</html>