package login

var List = map[string]string{"installation": `{{define "installation"}}
    <html>
    <head>
        <title>GoAdmin Install</title>
        <meta charset="utf-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, user-scalable=no">
        <link rel="stylesheet" href="../../assets/install/assets/css/main.css">
        <link rel="stylesheet" href="../../assets/fullpage/fullpage.min.css">
        <noscript>
            <link rel="stylesheet" href="../../assets/install/assets/css/noscript.css"/>
        </noscript>
    </head>
    <body class="">

    <div id="fullpage">
        <div class="section active" id="section1">
            <div class="slide" id="slide1">
                <div class="page-wrapper" style="margin-left: auto;margin-right:  auto;width: 1236px;">
                    <!-- Wrapper -->
                    <div class="wrapper" style="margin-left: auto;margin-right:  auto;">
                        <section class="panel color2-alt">
                            <div class="intro color2">
                                <h2 class="major">数据库配置</h2>
                                <p>目前仅支持Mysql</p>
                            </div>
                            <div class="inner columns aligned" style="padding: 2.5rem 3.5rem 2rem 3.5rem;">
                                <div class="span-2-5">
                                    <form method="post" action="#">
                                        <div class="field" style="width: calc(73% - 0.75rem);">
                                            <label for="demo-name">地址</label>
                                            <input type="text" name="demo-name" id="db-host" value=""
                                                   placeholder="127.0.0.1">
                                        </div>
                                        <div class="field quarter">
                                            <label for="demo-email">端口</label>
                                            <input type="email" name="demo-email" id="db-port" value=""
                                                   placeholder="3306">
                                        </div>
                                        <div class="field">
                                            <label for="demo-name">账号</label>
                                            <input type="text" name="demo-name" id="db-username" value=""
                                                   placeholder="root">
                                        </div>
                                        <div class="field">
                                            <label for="demo-email">密码</label>
                                            <input type="password" name="demo-email" id="db-password" value=""
                                                   placeholder="root">
                                        </div>
                                        <div class="field">
                                            <label for="demo-email">数据库名</label>
                                            <input type="text" name="demo-email" id="db-database" value=""
                                                   placeholder="database">
                                        </div>
                                        <div class="field" style="margin-top: 20px;">
                                            <ul class="actions">
                                                <li><input id="test-con-btn" type="submit" value="连接测试"
                                                           class="special color2"></li>
                                                <li><input class="input-next" type="reset" value="下一步"></li>
                                            </ul>
                                        </div>
                                    </form>
                                </div>
                            </div>
                        </section>
                        <!-- Copyright -->
                        <div class="copyright">
							<a target="_blank" href="https://beian.miit.gov.cn">黑ICP备-2023012515号</a>
                        </div>
                    </div>

                </div>
            </div>
            <div class="slide" id="slide2">
                <div class="page-wrapper" style="margin-left: auto;margin-right:  auto;width: 1236px;">
                    <!-- Wrapper -->
                    <div class="wrapper" style="margin-left: auto;margin-right:  auto;">
                        <section class="panel color2-alt">
                            <div class="intro color2">
                                <h2 class="major">选择管理数据表</h2>
                                <p>选择管理数据表</p>
                            </div>
                            <div class="inner columns aligned" style="padding: 2.5rem 3.5rem 2rem 3.5rem;">
                                <div class="span-2-5">
                                    <form method="post" action="#" id="table-div">
                                        <div class="field quarter">
                                            <input type="checkbox" id="demo-copy" name="demo-copy" class="color2">
                                            <label for="demo-copy">users</label>
                                        </div>
                                        <div class="field quarter">
                                            <input type="checkbox" id="demo-human" name="demo-human" class="color2"
                                                   checked="">
                                            <label for="demo-human">ninepic</label>
                                        </div>
                                        <div class="field quarter">
                                            <input type="checkbox" id="demo-copy1" name="demo-copy1" class="color2">
                                            <label for="demo-copy1">users</label>
                                        </div>
                                        <div class="field quarter">
                                            <input type="checkbox" id="demo-human1" name="demo-human1" class="color2"
                                                   checked="">
                                            <label for="demo-human1">ninepic</label>
                                        </div>
                                        <div class="field quarter">
                                            <input type="checkbox" id="demo-copy2" name="demo-copy2" class="color2">
                                            <label for="demo-copy2">users</label>
                                        </div>
                                        <div class="field quarter">
                                            <input type="checkbox" id="demo-human2" name="demo-human2" class="color2"
                                                   checked="">
                                            <label for="demo-human2">ninepic</label>
                                        </div>
                                        <div class="field quarter">
                                            <input type="checkbox" id="demo-copy3" name="demo-copy3" class="color2">
                                            <label for="demo-copy3">users</label>
                                        </div>
                                        <div class="field quarter">
                                            <input type="checkbox" id="demo-human3" name="demo-human3" class="color2"
                                                   checked="">
                                            <label for="demo-human3">ninepic</label>
                                        </div>
                                        <div class="field" style="margin-top: 20px;">
                                            <ul class="actions">
                                                <li><input class="input-prev" type="reset" value="上一步"></li>
                                                <li><input class="input-next" type="reset" value="下一步"></li>
                                            </ul>
                                        </div>
                                    </form>
                                </div>
                            </div>
                        </section>
                        <!-- Copyright -->
                        <div class="copyright">
							<a href="https://beian.miit.gov.cn">黑ICP备-2023012515号</a>
                        </div>
                    </div>
                </div>
            </div>
            <div class="slide" id="slide3">
                <div class="page-wrapper" style="margin-left: auto;margin-right:  auto;width: 1236px;">
                    <!-- Wrapper -->
                    <div class="wrapper" style="margin-left: auto;margin-right:  auto;">
                        <section class="panel color2-alt">
                            <div class="intro color2">
                                <h2 class="major">设置字段</h2>
                                <p>User表</p>
                            </div>
                            <div class="inner columns aligned"
                                 style="padding: 2.5rem 3.5rem 2rem 3.5rem;overflow-y: scroll;">
                                <div class="span-2-5">
                                    <form method="post" action="#">
                                        <div class="field" style="width: calc(38% - 0.5rem);">
                                            <label for="demo-name">id</label>
                                            <input type="text" name="demo-name" id="id" value="" placeholder="字段显示名称">
                                        </div>
                                        <div class="field third">
                                            <label for="demo-category">编辑表单类型</label>
                                            <div class="select-wrapper">
                                                <select name="demo-category" id="demo-category-id">
                                                    <option value="">-</option>
                                                    <option value="1">显示</option>
                                                    <option value="1">Text</option>
                                                    <option value="1">Select</option>
                                                </select>
                                            </div>
                                        </div>
                                        <div class="field third" style="margin-top: 40px;">
                                            <input type="checkbox" id="demo1-copy1" name="demo1-copy1" class="color2">
                                            <label for="demo1-copy1">取消</label>
                                        </div>

                                        <div class="field" style="width: calc(38% - 0.5rem);">
                                            <label for="demo-name">name</label>
                                            <input type="text" name="demo-name" id="name" value="" placeholder="name">
                                        </div>
                                        <div class="field third">
                                            <label for="demo-category">编辑表单类型</label>
                                            <div class="select-wrapper">
                                                <select name="demo-category" id="demo-category-name">
                                                    <option value="">-</option>
                                                    <option value="1">显示</option>
                                                    <option value="1">Text</option>
                                                    <option value="1">Select</option>
                                                </select>
                                            </div>
                                        </div>
                                        <div class="field third" style="margin-top: 40px;">
                                            <input type="checkbox" id="demo1-copy2" name="demo1-copy2" class="color2">
                                            <label for="demo1-copy2">取消</label>
                                        </div>

                                        <div class="field" style="width: calc(38% - 0.5rem);">
                                            <label for="demo-name">id</label>
                                            <input type="text" name="demo-name" value="" placeholder="id">
                                        </div>
                                        <div class="field third">
                                            <label for="demo-category">编辑表单类型</label>
                                            <div class="select-wrapper">
                                                <select name="demo-category">
                                                    <option value="">-</option>
                                                    <option value="1">显示</option>
                                                    <option value="1">Text</option>
                                                    <option value="1">Select</option>
                                                </select>
                                            </div>
                                        </div>
                                        <div class="field third" style="margin-top: 40px;">
                                            <input type="checkbox" id="demo1-copy3" name="demo1-copy3" class="color2">
                                            <label for="demo1-copy3">取消</label>
                                        </div>

                                        <div class="field" style="width: calc(38% - 0.5rem);">
                                            <label for="demo-name">name</label>
                                            <input type="text" name="demo-name" value="" placeholder="name">
                                        </div>
                                        <div class="field third">
                                            <label for="demo-category">编辑表单类型</label>
                                            <div class="select-wrapper">
                                                <select name="demo-category">
                                                    <option value="">-</option>
                                                    <option value="1">显示</option>
                                                    <option value="1">Text</option>
                                                    <option value="1">Select</option>
                                                </select>
                                            </div>
                                        </div>
                                        <div class="field third" style="margin-top: 40px;">
                                            <input type="checkbox" id="demo1-copy4" name="demo1-copy4" class="color2">
                                            <label for="demo1-copy4">取消</label>
                                        </div>

                                        <div class="field" style="width: calc(38% - 0.5rem);">
                                            <label for="demo-name">id</label>
                                            <input type="text" name="demo-name" value="" placeholder="id">
                                        </div>
                                        <div class="field third">
                                            <label for="demo-category">编辑表单类型</label>
                                            <div class="select-wrapper">
                                                <select name="demo-category">
                                                    <option value="">-</option>
                                                    <option value="1">显示</option>
                                                    <option value="1">Text</option>
                                                    <option value="1">Select</option>
                                                </select>
                                            </div>
                                        </div>
                                        <div class="field third" style="margin-top: 40px;">
                                            <input type="checkbox" id="demo1-copy5" name="demo1-copy5" class="color2">
                                            <label for="demo1-copy5">取消</label>
                                        </div>

                                        <div class="field" style="width: calc(38% - 0.5rem);">
                                            <label for="demo-name">name</label>
                                            <input type="text" name="demo-name" value="" placeholder="name">
                                        </div>
                                        <div class="field third">
                                            <label for="demo-category">编辑表单类型</label>
                                            <div class="select-wrapper">
                                                <select name="demo-category">
                                                    <option value="">-</option>
                                                    <option value="1">显示</option>
                                                    <option value="1">Text</option>
                                                    <option value="1">Select</option>
                                                </select>
                                            </div>
                                        </div>
                                        <div class="field third" style="margin-top: 40px;">
                                            <input type="checkbox" id="demo1-copy6" name="demo1-copy6" class="color2">
                                            <label for="demo1-copy6">取消</label>
                                        </div>

                                        <div class="field" style="width: calc(38% - 0.5rem);">
                                            <label for="demo-name">id</label>
                                            <input type="text" name="demo-name" value="" placeholder="id">
                                        </div>
                                        <div class="field third">
                                            <label for="demo-category">编辑表单类型</label>
                                            <div class="select-wrapper">
                                                <select name="demo-category">
                                                    <option value="">-</option>
                                                    <option value="1">显示</option>
                                                    <option value="1">Text</option>
                                                    <option value="1">Select</option>
                                                </select>
                                            </div>
                                        </div>
                                        <div class="field third" style="margin-top: 40px;">
                                            <input type="checkbox" id="demo1-copy7" name="demo1-copy7" class="color2">
                                            <label for="demo1-copy7">取消</label>
                                        </div>

                                        <div class="field" style="width: calc(38% - 0.5rem);">
                                            <label for="demo-name">name</label>
                                            <input type="text" name="demo-name" value="" placeholder="name">
                                        </div>
                                        <div class="field third">
                                            <label for="demo-category">编辑表单类型</label>
                                            <div class="select-wrapper">
                                                <select name="demo-category">
                                                    <option value="">-</option>
                                                    <option value="1">显示</option>
                                                    <option value="1">Text</option>
                                                    <option value="1">Select</option>
                                                </select>
                                            </div>
                                        </div>
                                        <div class="field third" style="margin-top: 40px;">
                                            <input type="checkbox" id="demo1-copy8" name="demo1-copy8" class="color2">
                                            <label for="demo1-copy8">取消</label>
                                        </div>

                                        <div class="field" style="width: calc(38% - 0.5rem);">
                                            <label for="demo-name">id</label>
                                            <input type="text" name="demo-name" value="" placeholder="id">
                                        </div>
                                        <div class="field third">
                                            <label for="demo-category">编辑表单类型</label>
                                            <div class="select-wrapper">
                                                <select name="demo-category">
                                                    <option value="">-</option>
                                                    <option value="1">显示</option>
                                                    <option value="1">Text</option>
                                                    <option value="1">Select</option>
                                                </select>
                                            </div>
                                        </div>
                                        <div class="field third" style="margin-top: 40px;">
                                            <input type="checkbox" id="demo1-copy9" name="demo1-copy9" class="color2">
                                            <label for="demo1-copy9">取消</label>
                                        </div>

                                        <div class="field" style="width: calc(38% - 0.5rem);">
                                            <label for="demo-name">name</label>
                                            <input type="text" name="demo-name" value="" placeholder="name">
                                        </div>
                                        <div class="field third">
                                            <label for="demo-category">编辑表单类型</label>
                                            <div class="select-wrapper">
                                                <select name="demo-category">
                                                    <option value="">-</option>
                                                    <option value="1">显示</option>
                                                    <option value="1">Text</option>
                                                    <option value="1">Select</option>
                                                </select>
                                            </div>
                                        </div>
                                        <div class="field third" style="margin-top: 40px;">
                                            <input type="checkbox" id="demo1-copy10" name="demo1-copy10" class="color2">
                                            <label for="demo1-copy10">取消</label>
                                        </div>

                                        <div class="field" style="width: calc(38% - 0.5rem);">
                                            <label for="demo-name">id</label>
                                            <input type="text" name="demo-name" value="" placeholder="id">
                                        </div>
                                        <div class="field third">
                                            <label for="demo-category">编辑表单类型</label>
                                            <div class="select-wrapper">
                                                <select name="demo-category">
                                                    <option value="">-</option>
                                                    <option value="1">显示</option>
                                                    <option value="1">Text</option>
                                                    <option value="1">Select</option>
                                                </select>
                                            </div>
                                        </div>
                                        <div class="field third" style="margin-top: 40px;">
                                            <input type="checkbox" id="demo1-copy11" name="demo1-copy11" class="color2">
                                            <label for="demo1-copy11">取消</label>
                                        </div>

                                        <div class="field" style="width: calc(38% - 0.5rem);">
                                            <label for="demo-name">name</label>
                                            <input type="text" name="demo-name" value="" placeholder="name">
                                        </div>
                                        <div class="field third">
                                            <label for="demo-category">编辑表单类型</label>
                                            <div class="select-wrapper">
                                                <select name="demo-category">
                                                    <option value="">-</option>
                                                    <option value="1">显示</option>
                                                    <option value="1">Text</option>
                                                    <option value="1">Select</option>
                                                </select>
                                            </div>
                                        </div>
                                        <div class="field third" style="margin-top: 40px;">
                                            <input type="checkbox" id="demo1-copy12" name="demo1-copy12" class="color2">
                                            <label for="demo1-copy12">取消</label>
                                        </div>

                                        <div class="field" style="margin-top: 20px;">
                                            <ul class="actions">
                                                <li><input class="input-prev" type="reset" value="上一步"></li>
                                                <li><input class="input-next" type="reset" value="下一步"></li>
                                            </ul>
                                        </div>
                                    </form>
                                </div>
                            </div>
                        </section>
                        <!-- Copyright -->
                        <div class="copyright">
							<a target="_blank" href="https://beian.miit.gov.cn">黑ICP备-2023012515号</a>
                        </div>
                    </div>
                </div>
            </div>
            <div class="slide" id="slide4">
                <div class="page-wrapper" style="margin-left: auto;margin-right:  auto;width: 1236px;">
                    <!-- Wrapper -->
                    <div class="wrapper" style="margin-left: auto;margin-right:  auto;">
                        <section class="panel color2-alt">
                            <div class="intro color2">
                                <h2 class="major">设置超级管理员</h2>
                                <p>超级管理员设置</p>
                            </div>
                            <div class="inner columns aligned" style="padding: 2.5rem 3.5rem 2rem 3.5rem;">
                                <div class="span-2-5">
                                    <form method="post" action="#">
                                        <div class="field">
                                            <label for="demo-name">账号</label>
                                            <input type="text" name="demo-name" id="username" value=""
                                                   placeholder="root">
                                        </div>
                                        <div class="field">
                                            <label for="demo-email">密码</label>
                                            <input type="password" name="demo-email" id="password" value=""
                                                   placeholder="root">
                                        </div>
                                        <div class="field">
                                            <label for="demo-email">确认密码</label>
                                            <input type="password" name="demo-password_comfirm" id="password_comfirm"
                                                   value=""
                                                   placeholder="root">
                                        </div>
                                        <div class="field" style="margin-top: 20px;">
                                            <ul class="actions">
                                                <li><input class="input-next" type="reset" value="Enjoy!"></li>
                                            </ul>
                                        </div>
                                    </form>
                                </div>
                            </div>
                        </section>
                        <!-- Copyright -->
                        <div class="copyright">
                        </div>
                    </div>

                </div>
            </div>
        </div>
    </div>

    <!-- Scripts -->
    <script src="../../assets/install/assets/js/jquery.min.js"></script>
    <script src="../../assets/install/assets/js/skel.min.js"></script>
    <script src="../../assets/install/assets/js/main.js"></script>
    <script src="../../assets/fullpage/fullpage.min.js"></script>
    <script type="text/javascript">
        var myFullpage = new fullpage('#fullpage', {
            anchors: ['first'],
            lazyLoad: true,
            scrollBar: true,
            autoScrolling: false
        });

        myFullpage.destroy();

        $(".input-next").click(function (e) {
            myFullpage.moveSlideRight()
        });
        $(".input-prev").click(function (e) {
            myFullpage.moveSlideLeft()
        });

        /** 处理逻辑 **/

        // 测试连接
        $('#test-con-btn').on('click', function (e) {
            e.preventDefault();
            $.ajax({
                dataType: 'json',
                type: 'POST',
                url: '/install/database/check',
                async: 'true',
                data: {
                    'h': $("#db-host").val(),
                    'po': $("#db-port").val(),
                    'u': $("#db-username").val(),
                    'pa': $("#db-password").val(),
                    'db': $("#db-database").val()
                },
                success: function (data) {
                    console.log(data)
                    if (data.code === 0) {

                        // html = "";
                        //
                        // for (i = 0; i < data.data.list.length; i++) {
                        //     html += '<div class="field"><input type="checkbox" id="' +
                        //             data.data.list[i] +
                        //             '" name="table-' + i + '" class="color2"><label for="table-' + i + '">' +
                        //             data.data.list[i]
                        //             + '</label></div>';
                        // }
                        //
                        // $("#table-div").html(html);

                        alert(data.msg);

                    } else {
                        alert(data.msg);
                    }
                },
                error: function (data) {
                    alert("ok");
                }
            });
        });

    </script>
    <style>
        .fp-controlArrow {
            display: none;
        }
    </style>
    </body>
    </html>
{{end}}`}
