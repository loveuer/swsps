<template>
    <div class="container">
        <div class="login-zone">
            <div class="login">
                <div class="form">
                    <form>
                        <div class="input-zone">
                            <label for="username">用户名:</label><br>
                            <span class="input-logo"></span>
                            <span>
                                <input type="text" name="username">
                            </span>
                            <span class="input-warning"></span>
                        </div>
                        <div class="input-zone">
                            <label for="password">密码:</label><br>
                            <span class="input-logo"></span>
                            <span>
                                <input type="password" name="password">
                            </span>
                            <span class="input-warning"></span>
                        </div>

                        <div class=check-zone>
                            <input type="checkbox" id="remember"><label for="remember">Remember me</label>
                            <a href="/forget">Forgot password</a>
                        </div>
                        <div class="submit-zome">
                            <button type="button">Log in</button>
                        </div>
                        <div class="register-zone">
                            or <a href="/register">register</a> now!
                        </div>
                    </form>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
export default {
    data() {
        return {
            formLayout: "horizontal",
        };
    },
    mounted() {

    },
    methods: {
        handleSubmit: () => {
            let ips = document.querySelectorAll("input");
            let username = ips[0];
            let password = ips[1];
            if(username === ""){
                username.style.border = "1px solid red";
            }
        },
    }
};
</script>

<style scoped>
.container{
    position: absolute;
    height: 100%;
    width: 100%;
    background-image: url("~@/assets/img/runway.png");
    background-repeat: no-repeat;
    background-position: center;
    background-size: cover;
}
.login-zone{
    height: 100%;
    width: 100%;
    background-color: rgba(0,0,0,0.3);
    filter:Alpha(opacity=50);
    display: flex;
    justify-content: center;
    align-items: center;
}
.login{
    height: 250px;
    width: 340px;
    border: 2px solid #fff;
    padding: 20px;
    border-radius: 20px;
}
.login:hover{
    background-color: rgba(255,255,255,0.4);
}
#submit{
    margin-top: 20px;
}
</style>
