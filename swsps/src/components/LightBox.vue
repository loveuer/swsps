<template>
    <div id="container" @click="out_closelb" @keypress.esc="console.log('esc pressed')">
        <div class="frame">
            <div class="img-zone">
                <div class="close-btn" @click="closelb">
                    <i class="el-icon-close"></i>
                </div>
                <img :src="lightboxImg.img" @click="openImg" id="lb-img">
            </div>
            <div class="alt-zone">
                {{ lightboxImg.alt }}
            </div>
        </div>
    </div>
</template>

<script>
import lb from "lightbox2";
import jq from "jquery";

export default {
    data() {
        return {

        };
    },
    props: {
        lightboxImg: Object,
    },
    methods: {
        openImg: function() {
            window.open(this.lightboxImg.img, '_blank');
        },
        closelb: function() {
            this.$emit('lbclose');
        },
        out_closelb: function() {
            console.log('out-close');
            this.closelb();
        },
    },
    computed: {

    },
    mounted() {
        window.$ = jq;
    },
    beforeDestroy() {
        window.$ = null;
    },
}
</script>

<style scoped>
#container {
    height: 100%;
    min-height: 100%;
    width: 100%;
    background-color: rgba(0,0,0,0.8);
    position: absolute;
    top: 0;
    left: 0;
    display: flex;
    align-items: center;
    justify-content: center;
}
.frame {
    max-width: 75%;
    width: auto;
    height: 85%;
}
.img-zone {
    width: 100%;
    height: 100%;
    background-color: rgba(0,0,0,0.4);
    border: 2px solid #eee;
    border-radius: 10px;
    position: relative;
    cursor: zoom-in;
    box-sizing: border-box;
}
.close-btn {
    position: absolute;
    top: -10px;
    right: -10px;
    width: 30px;
    height: 30px;
    background-color: #fff;
    border-radius: 50%;
    display: flex;
    align-items: center;
    justify-content: center;
    cursor: pointer;
    outline: none;
    border: 2px solid #eee;
}
.alt-zone {
    width: 100%;
    height: 40px;
    line-height: 40px;
    color: #ddd;
    text-align: right;
}
.img-zone > img {
    height: 100%;
    width: auto;
    border-radius: 8px;
}
</style>

