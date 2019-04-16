<template>
    <div id="container">
        <div>
            <u-menu></u-menu>
        </div>
        <div class="photos">
            <div>
                <el-carousel indicator-position="outside" height="800px">
                    <el-carousel-item v-for="(p, index) in photos" :key="index">
                        <h3>
                            <div v-bind:style="{backgroundImage:'url('+p+')'}">
                                <img />
                            </div>
                        </h3>
                    </el-carousel-item>
                </el-carousel>
            </div>
        </div>
    </div>
</template>

<script>
import uMenu from "./u-cmpnts/u_menu.vue";

export default {
    data() {
        return {
            photos: [
                "http://localhost:8000/static/img/frame_photo_1.jpg",
                "http://localhost:8000/static/img/frame_photo_2.jpg",
                "http://localhost:8000/static/img/frame_photo_3.jpg",
                "http://localhost:8000/static/img/frame_photo_4.jpg",
            ],
        };
    },
    components: {
        "u-menu": uMenu,
    },
};
</script>

<style scoped>
#container {
    width: 100%;
}
.photos {
    width: 100%;
    display: flex;
    justify-content: center;
    margin-top: 30px;
}
.photos > div {
    width: 90%;
}
.el-carousel__item h3 {
    color: #475669;
    font-size: 18px;
    opacity: 0.75;
    line-height: 800px;
    margin: 0;
}
.el-carousel__item h3 > div {
    background-repeat: no-repeat;
    background-size: cover;
    background-position: center;
}
.el-carousel__item h3 > div > img {
    height: 100%;
    width: 100%;
}
.el-carousel__item:nth-child(2n) {
    background-color: #99a9bf;
}

.el-carousel__item:nth-child(2n+1) {
    background-color: #d3dce6;
}
</style>
