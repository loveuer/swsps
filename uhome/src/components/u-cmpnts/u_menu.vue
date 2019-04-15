<template>
    <div id="container">
        <div>balabala</div>
        <div>balabala</div>
        <div></div>
        <div>balabala</div>
        <div>balabala</div>
    </div>
</template>

<script>
export default {
    
};
</script>

<style scoped>
#container {
    height: 50px;
    width: 100%;
    position: fixed;
    top: 0;
    left: 0;
    background-color: aquamarine;
    display: flex;
    justify-content: center;
    align-items: center;
}
#container > div {
    width: 100px;
    height: 100%;
    margin: 10px;
}
#container > div:nth-child(3) {
    background-image: url("@~/assets/img/uhome.png");
    background-repeat: no-repeat;
    background-size: contain;
    background-position: center;
}
</style>
