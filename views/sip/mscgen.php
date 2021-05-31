<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>Mscgen diagram</title>
    <script src="/public/js/mscgen-inpage.js" defer></script>
    <link href="/public/bootstrap/css/bootstrap.min.css" rel="stylesheet" />
    <script type="text/javascript" src="/public/js/axios.min.js"></script>
    <link href="/public/css/font-awesome.min.css" rel="stylesheet" />
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
    <div id="box">
        <div id="left">
            <pre id="mscremote" class="mscgen_js" style="max-height: 850px;">
                msc {
                <?php
                    foreach ($sip_dia_array as $value) {
                        echo $value;
                    }
                ?>
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
    var path = "<?php echo $path; ?>"
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
    function show(a) {
        var pDiv = document.getElementById("showarea")
        axios.get("/index.php/Review/return_detail_sip?p=" + path)
        .then(function(response) {
            console.log(response.data[a])
            pDiv.innerHTML = response.data[a].desc
        })
        .catch(function(error) {
            console.log(error)
        })
    }
</script>

</html>
