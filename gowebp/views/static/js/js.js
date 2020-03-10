function adduser() {
    var element = document.getElementsByClassName("content")[0];
    element.innerHTML = `
    <form action="/AddUser" method="post">
        <legend>添加用户</legend>
        <label for="name">用户名：</label>
        <input type="text" name="name" id="name" required><br>
        <label for="password">密&nbsp;&nbsp;&nbsp;&nbsp;码：</label>
        <input type="password" name="password" id="password" required><br>
        <label for="email">邮&nbsp;&nbsp;&nbsp;&nbsp;箱：</label>
        <input type="email" name="email" id="email"><br>
        <input type="submit" value="添加">
        <input type="reset" value="清空">
    </form>`;
}

function adduserv2() {
    var element = document.getElementsByClassName("content")[0];
    element.innerHTML = `
    <fieldset>
        <legend>添加用户</legend>
        <label for="name">用户名：</label>
        <input type="text" name="name" id="name" required><br>
        <label for="password">密&nbsp;&nbsp;&nbsp;&nbsp;码：</label>
        <input type="password" name="password" id="password" required><br>
        <label for="email">邮&nbsp;&nbsp;&nbsp;&nbsp;箱：</label>
        <input type="email" name="email" id="email"><br>
        <input type="button" value="添加" onclick="ajaxadduser()">
        <input type="reset" value="清空">
    </fieldset>`;
}

function ajaxadduser() {
    var request = new XMLHttpRequest();
    request.open("post", "/AddUser");

    request.setRequestHeader("Content-Type", "application/x-www-form-urlencoded");
    request.send("name=" + document.getElementById("name").value + "&password=" + document.getElementById("password").value + "&email=" + document.getElementById("email").value);

    request.onreadystatechange = function () {
        if (request.readyState == 4 && request.status == 200) {
            var result = request.responseText;
            window.alert(result);
            document.getElementById("name").value = "";
            document.getElementById("password").value = "";
            document.getElementById("email").value = "";
            return
        }
        if (request.readyState == 4 && request.status == 500) {
            var result = request.responseText
            window.alert(result)
        }
    }

}

function deleteuser() {
    var element = document.getElementsByClassName("content")[0];
    element.innerHTML = `
    <form action="/DeleteUser" method="post">
        <legend>删除用户</legend>
        <label for="name">用户名：</label>
        <input type="text" name="name" id="name" required><br>
        <label for="password">密&nbsp;&nbsp;&nbsp;&nbsp;码：</label>
        <input type="password" name="password" id="password" required><br>
        <label for="email">邮&nbsp;&nbsp;&nbsp;&nbsp;箱：</label>
        <input type="email" name="email" id="email"><br>
        <input type="submit" value="删除">
        <input type="reset" value="清空">
    </form>`;
}
function modifyuser() {
    var element = document.getElementsByClassName("content")[0];
    element.innerHTML = `
    <form action="/ModifyUser" method="post">
        <legend>修改用户</legend>
        <label for="name">用户名：</label>
        <input type="text" name="name" id="name" required><br>
        <label for="password">密&nbsp;&nbsp;&nbsp;&nbsp;码：</label>
        <input type="password" name="password" id="password" required><br>
        <label for="email">邮&nbsp;&nbsp;&nbsp;&nbsp;箱：</label>
        <input type="email" name="email" id="email" required><br>
        <input type="submit" value="提交">
        <input type="reset" value="清空">
    </form>`;
}

function queryuser() {
    var element = document.getElementsByClassName("content")[0];
    element.innerHTML = `
    <form action="/QueryUser" method="get">
        <legend>查看用户</legend>
        <label for="name">用户名：</label>
        <input type="text" name="name" id="name"><br>
        <label for="password">密&nbsp;&nbsp;&nbsp;&nbsp;码：</label>
        <input type="password" name="password" id="password"><br>
        <label for="email">邮&nbsp;&nbsp;&nbsp;&nbsp;箱：</label>
        <input type="email" name="email" id="email"><br>
        <input type="submit" value="查找">
        <input type="reset" value="清空">
    </form>`;
}

function queryuserv2() {
    var element = document.getElementsByClassName("content")[0];
    element.innerHTML = `
    <form action="/QueryUser" method="get">
        <legend>查看用户</legend>
        <label for="name">用户名：</label>
        <input type="text" name="name" id="name"><br>
        <label for="password">密&nbsp;&nbsp;&nbsp;&nbsp;码：</label>
        <input type="password" name="password" id="password"><br>
        <label for="email">邮&nbsp;&nbsp;&nbsp;&nbsp;箱：</label>
        <input type="email" name="email" id="email"><br>
        <input type="button" value="查找" onclick="ajaxqueryuser()">
        <input type="reset" value="清空">
    </form>`;
}

function ajaxqueryuser() {
    var request = new XMLHttpRequest()
    request.open("get", "/QueryUser?name="+ document.getElementById("name").value)
    request.send()

    request.onreadystatechange = function () {
        if (request.readyState == 4 && request.status == 200) {
            var userjson = JSON.parse(request.responseText)

            var content = document.getElementsByClassName("result")[0];

            var htmlStr = `<table border="1px"><caption>用户信息</caption>
                    <tr><th>用户ID</th><th>用户姓名</th><th>用户邮箱</th></tr>`;
            htmlStr += `<tr><td>` + userjson.ID + `</td><td>` + userjson.UserName + `</td><td>` + userjson.Email + `</td></tr>`;
            htmlStr += `</table>`;
            content.innerHTML = htmlStr;
        }
    }

}


function listuser() {
    var element = document.getElementsByClassName("content")[0];
    element.innerHTML = `
    <form action="/ListUser" method="get">
        <legend>列出用户</legend>
        <input type="submit" value="查找">
    </form>`;
}

function listuserv2() {
    document.getElementsByClassName("result")[0].innerHTML = "";
    var element = document.getElementsByClassName("content")[0];
    element.innerHTML = `
    <fieldset>
        <legend>列出用户</legend>
        <input type="button" value="查找" onclick="ajaxlistuser()">
    </fieldset>`;
}

function ajaxlistuser() {
    var request = new XMLHttpRequest();
    request.open("get", "http://localhost:8080/ListUser");
    request.send();

    request.onreadystatechange = function () {
        if (request.readyState == 4 && request.status == 200) {
            var users = JSON.parse(request.responseText);
            var content = document.getElementsByClassName("result")[0];

            var htmlStr = `<table border="1px"><caption>用户信息</caption>
            <tr><th>用户ID</th><th>用户姓名</th><th>用户邮箱</th></tr>`;
            for (let index = 0; index < users.length; index++) {
                htmlStr += `<tr><td>` + users[index].ID + `</td><td>` + users[index].UserName + `</td><td>` + users[index].Email + `</td></tr>`;
            }
            htmlStr += `</table>`;
            content.innerHTML = htmlStr;
        }
    }
}