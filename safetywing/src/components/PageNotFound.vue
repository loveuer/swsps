<template>
    <div class="container">
        <div class="logo">
            <div onclick="window.location.href='/'"></div>
        </div>
        <div class="reminder">
            <div></div>
            <div>404 Page Not Found</div>
        </div>
        <div class="jump">
            <div id="leftTime"></div>
            <div>秒后返回主页</div>
            <div onclick="window.location.href='/'">立即返回</div>
        </div>
    </div>
</template>

<script>
function setCountDownTimer() {
    let countDown = 5;
    document.querySelector("#leftTime").innerHTML = countDown;
    let interval = setInterval(() => {
        countDown--;
        document.querySelector("#leftTime").innerHTML = countDown;

        if(countDown === 0){
            clearInterval(interval);
            window.location.href= "/";
        }
    }, 1000);
};

export default {
    mounted() {
        setCountDownTimer();
    },
}
</script>


<style scoped>
.container{
    display: flex;
    flex-wrap: wrap;
    justify-content: center;
    width: 100%;
}
.logo{
    width: 100%;
    height: 200px;
    display: flex;
    justify-content: center;
}
.logo > div {
    height: 100%;
    width: 350px;
    background-image: url("~@/assets/img/logo.png");
    background-repeat: no-repeat;
    background-position: center;
    background-size: contain;
    cursor: pointer;
}
.reminder{
    height: 100px;
    width: 100%;
    display: flex;
    justify-content: center;
    align-items: center;
}
.reminder > div:first-child {
    width: 100px;
    height: 50%;
    background-image: url("~@/assets/img/404.svg");
    background-repeat: no-repeat;
    background-position: center;
    background-size: contain;
}
.reminder > div:nth-child(2) {
    font-size: 24px;
    color: #0082E4;
}
.jump{
    height: 100px;
    width: 100%;
    display: flex;
    justify-content: center;
    align-items: center;
}
.jump > div:nth-child(3){
    color: #0082E4;
    margin-left: 20px;
    cursor: pointer;
}
.jump > div:nth-child(3):hover{
    text-decoration: #0082E4 underline;
}
#leftTime{
    color: #ED4337;
}
</style>
