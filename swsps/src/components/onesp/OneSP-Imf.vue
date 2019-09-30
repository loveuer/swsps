<template>
    <div id="container">
        <div class="imf-form">
            <div class="onesp-row">
                <div>P/N</div>
                <div>{{ imfsp.pn }}</div>
            </div>
            <div class="onesp-row">
                <div>S/N</div>
                <div>{{ imfsp.sn }}</div>
            </div>
            <div class="onesp-row">
                <div>原模拟机</div>
                <div>{{ imfsp.orgsim }}</div>
            </div>
            <div class="onesp-row">
                <div>现模拟机</div>
                <div>{{ imfsp.nowsim }}</div>
            </div>
            <div class="onesp-row">
                <div>位置</div>
                <div>{{ imfsp.pos }}</div>
            </div>
            <div class="onesp-row">
                <div>状态</div>
                <div>{{ imfsp.status }}</div>
            </div>
            <div class="onesp-row">
                <div>数量</div>
                <div>{{ imfsp.amount }}</div>
            </div>
            <div class="row-comment">
                <div>备注</div>
                <div v-html="imfsp.comment.replace(/\r?\n/g,'<br/>')"></div>
            </div>
            <div class="onesp-imgs">
                <div>图片</div>
                <div class="imgs-zone">
                    <div class="one-img">
                        <!-- <img :src="imfsp.img1" @click="handle_lb('img1')"> -->
                        <a :href="imfsp.img1" data-lightbox="onesp-image" data-title="image-1">
                            <img :src="imfsp.img1">
                        </a>
                    </div>
                    <div class="one-img">
                        <!-- <img :src="imfsp.img2" @click="handle_lb('img2')"> -->
                        <a :href="imfsp.img2" data-lightbox="onesp-image" data-title="image-2">
                            <img :src="imfsp.img2">
                        </a>
                    </div>
                </div>
            </div>
            <div style="display:flex;margin-top:30px;">
                <el-button type="primary" @click="editsp">编辑</el-button>
                <el-button type="primary" @click="showsphis">历史</el-button>
            </div>
        </div>
        <div v-if="show_lightbox">
            <light-box :lightboxImg="lbimg" v-on:lbclose="close_lightbox" ref="lbimg"></light-box>
        </div>
    </div>
</template>

<script>
require ("lightbox2");
import "lightbox2/dist/css/lightbox.min.css";
import jq from "jquery";

export default {
    data() {
        return {
            show_lightbox: false,
            lbimg: {},
        };
    },
    props: {
        onesp: Object,
    },
    methods: {
        showsphis: function() {
            this.$emit('chgmod', 'sp-his');
        },
        handle_lb: function(which) {
            this.lbimg.img = this.imfsp[which];
            this.lbimg.alt = this.imfsp.name;
            this.show_lightbox = true;
        },
        close_lightbox: function() {
            this.show_lightbox = false;
        },
        try_closelb: function() {
            this.show_lightbox = false;
        },
        editsp: function() {
            this.$emit("chgmod", 'sp-edit');
        },
    },
    computed: {
        imfsp: function() {
            let deepcopysp = JSON.parse(JSON.stringify(this.onesp));
            if (deepcopysp.img1 === 'None') {
                // deepcopysp.img1 = 'http://192.168.0.19:8820/static/sp_img/None.jpg';
                deepcopysp.img1 = `http://${window.location.host}/static/sp_img/None.jpg`;
            };
            if (deepcopysp.img2 === 'None') {
                // deepcopysp.img2 = 'http://192.168.0.19:8820/static/sp_img/None.jpg';
                deepcopysp.img2 = `http://${window.location.host}/static/sp_img/None.jpg`;
            };
            return deepcopysp;
        }
    },
    created() {
        window.$ = jq;
    },
    beforeDestroy() {
        window.$ = null;
    },
    components: {
        // 'light-box': lightbox,
    },
};
</script>

<style scoped>
@font-face {
    font-family: hack;
    src: url("~@/assets/fonts/Hack-BoldItalic.ttf");
}
#container {
    width: 100%;
    display: flex;
    justify-content: center;
}
.imf-form {
    width: 80%;
    margin-top: 50px;
    margin-bottom: 100px;
}
.onesp-row {
    width: 100%;
    display: flex;
    height: 50px;
    line-height: 50px;  
    border-bottom: 1px solid #e8e8e8;
}
.onesp-row:hover {
    background-color: #e8e8e8;
    border-bottom: 1px solid rgba(131, 208, 242, 1);
}
.onesp-row > div:first-child {
    width: 150px;
    color: #666;
    text-indent: 5px;
    font-size: 0.85rem;
    font-family: hack;
    font-weight: bold;
}
.row-comment {
    width: 100%;
    display: flex;
    line-height: 50px;
    min-height: 50px;
    border-bottom: 1px solid #e8e8e8;
}
.row-comment:hover {
    background-color: #e8e8e8;
    border-bottom: 1px solid rgba(131, 208, 242, 1);
}
.row-comment > div:first-child {
     width: 150px;
    color: #666;
    text-indent: 5px;
    font-size: 0.85rem;
    font-family: hack;
    font-weight: bold;
}
.onesp-imgs {
    width: 100%;
    height: 200px;
    line-height: 200px;
    display: flex;
}

.onesp-imgs > div:first-child {
    width: 150px;
    color: #666;
    text-indent: 5px;
    font-size: 0.85rem;
    font-family: hack;
    font-weight: bold;
}
.imgs-zone {
    width: 100%;
    height: 100%;
    display: flex;
    align-items: center;
}
.one-img {
    height: 180px;
    border: 1px solid #fff;
    border-radius: 5px;
}
.one-img:first-child{
    margin-left: 10px;
}
.one-img:last-child{
    margin-left: 20px;
}
.one-img:hover {
    border: 1px solid rgba(131, 208, 242, 1);
    box-shadow: 0 0 5px rgba(131, 208, 242, 1);
}
.one-img img {
    height: 180px;
    width: auto;
    border-radius: 5px;
    cursor: pointer;
}
</style>
