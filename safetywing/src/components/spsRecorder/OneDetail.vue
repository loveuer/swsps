<template>
    <div style="width:80%;">
        <div class="one-name">
            <div></div>
            <div>
                <font>{{ sp.name }}</font>
            </div>
        </div>
        <div>
            <table>
                <tr v-if="this.$store.state.user.is_admin > 1">
                    <th>ID</th>
                    <td>{{ sp.id }}</td>
                </tr>
                <tr>
                    <th>P/N</th>
                    <td>{{ sp.pn }}</td>
                </tr>
                <tr>
                    <th>S/N</th>
                    <td>{{ sp.sn }}</td>
                </tr>
                <tr>
                    <th>原模拟机</th>
                    <td>{{ sp.orgsim }}</td>
                </tr>
                <tr>
                    <th>模拟机</th>
                    <td>{{ sp.nowsim }}</td>
                </tr>
                <tr>
                    <th>位置</th>
                    <td>{{ sp.pos }}</td>
                </tr>
                <tr>
                    <th>状态</th>
                    <td>{{ sp.status }}</td>
                </tr>
                <tr>
                    <th>数量</th>
                    <td>{{ sp.amount }}</td>
                </tr>
                <tr>
                    <th>备注</th>
                    <td>{{ sp.comment }}</td>
                </tr>
            </table>
        </div>
        <div class="one-imgs">
            <!-- <el-col :span="6"> -->
            <div>
                <div class="imgs-frame">
                    <div>
                        <img 
                            :src="sp.img1" 
                            @error="imgsrc2none('img1')" 
                            height="180" 
                            width="auto"
                            @click="showLightBox('img1')">
                    </div>
                </div>
            </div>
            <!-- </el-col> -->
            <!-- <el-col :span="6"> -->
            <div style="margin-left:50px;">
                <div class="imgs-frame">
                    <div>
                        <img :src="sp.img2" @error="imgsrc2none('img2')" height="180" width="auto" @click="showLightBox('img1')">
                    </div>
                </div>
            </div>
            <!-- </el-col> -->
        </div>
        <div class="light-box">
            <loveuer-lightbox ref="uLightBox"></loveuer-lightbox>
        </div>
    </div>
</template>

<script>
import loveuerLightBox from '../uLightBox.vue';

export default {
    data() {
        return {
            sp: {},
            lightBoxDisplay: 'none',
        }
    },
    methods: {
        imgsrc2none: function(img) {
            this.sp[img] = "http://localhost:8000/static/img/spimg/None.jpg";
        },
        showLightBox: function(img) {
            this.$refs.uLightBox.ifshow = true;
            this.$refs.uLightBox.imgSrc = this.sp[img];
        },
    },
    mounted() {
        this.$http.get("/api/sps/one/" + this.$route.params.spid)
            .then(resp => {
                this.sp = resp.data;
            })
            .catch(error => { console.log("sps detail http error: \n", error.reponse); });
    },
    components: {
        "loveuer-lightbox": loveuerLightBox,
    },
};
</script>

<style scoped>
.one-name {
    height: 50px;
    width: 100%;
    display: flex;
    background-image: url("~@/assets/img/one-name.png");
    background-repeat: no-repeat;
    background-size: contain;
    background-position: left 0 top 0;
}
.one-name > div:first-child {
    height: 100%;
    width: 200px;
}
.one-name > div:last-child{
    width: 100%;
    border-bottom: 8px solid #409EFF;
}
@font-face {
    font-family: "back_bold_italic";
    src: url("~@/assets/font/Hack-BoldItalic.ttf");
}
.one-name > div > font {
    font-size: 30px;
    text-indent: -20px;
    font-family: "back_bold_italic";
}
.one-imgs {
    margin-top: 20px;
    height: 300px;
    width: 100%;
    display: flex;
}
.imgs-frame {
    height: 200px;
    display: flex;
    align-items: center;
    justify-content: center;
}
.imgs-frame > div {
    height: 180px;
    border: 1px solid #ddd;
    border-radius: 10px;
    cursor: pointer;
}
.imgs-frame > div:hover {
    border: 1px solid #409EFF;
    box-shadow: 0 0 5px #409EFF;
}
.imgs-frame > div > img {
    border-radius: 10px;
}
table {
    margin-top: 15px;
}
table > tr {
    height: 50px;
    border-bottom: 1px solid #ddd;
}
table > tr > th {
    width:165px;
    text-align: left;
}
</style>
