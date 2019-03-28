<template>
    <div id="training-recoder">
        <div style="width:100%;">
            <loveuer-menu ref="uMenu"></loveuer-menu>
        </div>
        <div style="margin-top:30px;margin-left:30px;">
            <el-form ref="form" :model="form" label-width="100px">
                <el-row>
                    <el-col :span="8">
                    <el-form-item label="训练日期:">
                        <el-date-picker type="date" placeholder="选择日期" v-model="form.date" style="width: 100%;" @change="showPickedDate"></el-date-picker>
                    </el-form-item>

                    <el-form-item label="模拟机编号:">
                        <el-select v-model="form.sim" placeholder="" style="width: 100%;">
                            <el-option label="5978" value="5978"></el-option>
                            <el-option label="5989" value="5989"></el-option>
                            <el-option label="5008" value="5008"></el-option>
                            <el-option label="5015" value="5015"></el-option>
                        </el-select>
                    </el-form-item>

                    <el-form-item label="航空公司:">
                        <el-select v-model="form.airline" placeholder="" style="width: 100%;">
                            <el-option label="国　　航" value="国　　航"></el-option>
                            <el-option label="藏　　航" value="藏　　航"></el-option>
                            <el-option label="重庆航空" value="重庆航空"></el-option>
                            <el-option label="成都航空" value="成都航空"></el-option>
                            <el-option label="东方航空" value="东方航空"></el-option>
                            <el-option label="试飞中心" value="试飞中心"></el-option>
                            <el-option label="华夏航空" value="华夏航空"></el-option>
                            <el-option label="藏航理论" value="藏航理论"></el-option>
                            <el-option label="天府航空" value="天府航空"></el-option>
                            <el-option label="其　　他" value="其　　他"></el-option>
                        </el-select>
                    </el-form-item>

                    <el-form-item label="训练场次:">
                        <el-select v-model="form.session" placeholder="" style="width: 100%;">
                            <el-option label="第 1 场" value="1"></el-option>
                            <el-option label="第 2 场" value="2"></el-option>
                            <el-option label="第 3 场" value="3"></el-option>
                            <el-option label="第 4 场" value="4"></el-option>
                            <el-option label="第 5 场" value="5"></el-option>
                            <el-option label="第 6 场" value="6"></el-option>
                            <el-option label="第 7 场" value="7"></el-option>
                            <el-option label="第 8 场" value="8"></el-option>
                            <el-option label="第 9 场" value="9"></el-option>
                        </el-select>
                    </el-form-item>

                    <el-form-item label="训练时间:">
                        <el-col :span="8">
                            <el-select v-model="form.start" placeholder="" style="width: 100%;" @change="calTotalTime">
                                <el-option label="00 : 00" value="0"></el-option>
                                <el-option label="01 : 00" value="1"></el-option>
                                <el-option label="02 : 00" value="2"></el-option>
                                <el-option label="03 : 00" value="3"></el-option>
                                <el-option label="04 : 00" value="4"></el-option>
                                <el-option label="05 : 00" value="5"></el-option>
                                <el-option label="06 : 00" value="6"></el-option>
                                <el-option label="07 : 00" value="7"></el-option>
                                <el-option label="08 : 00" value="8"></el-option>
                                <el-option label="09 : 00" value="9"></el-option>
                                <el-option label="10 : 00" value="10"></el-option>
                                <el-option label="11 : 00" value="11"></el-option>
                                <el-option label="12 : 00" value="12"></el-option>
                                <el-option label="13 : 00" value="13"></el-option>
                                <el-option label="14 : 00" value="14"></el-option>
                                <el-option label="15 : 00" value="15"></el-option>
                                <el-option label="16 : 00" value="16"></el-option>
                                <el-option label="17 : 00" value="17"></el-option>
                                <el-option label="18 : 00" value="18"></el-option>
                                <el-option label="19 : 00" value="19"></el-option>
                                <el-option label="20 : 00" value="20"></el-option>
                                <el-option label="21 : 00" value="21"></el-option>
                                <el-option label="22 : 00" value="22"></el-option>
                                <el-option label="23 : 00" value="23"></el-option>
                                <el-option label="24 : 00" value="24"></el-option>
                            </el-select>
                        </el-col>
                        <el-col class="line" :span="2" style="text-align:center;">-</el-col>
                        <el-col :span="8">
                            <el-select v-model="form.end" placeholder="" style="width: 100%;"  @change="calTotalTime">
                                <el-option label="00 : 00" value="0"></el-option>
                                <el-option label="01 : 00" value="1"></el-option>
                                <el-option label="02 : 00" value="2"></el-option>
                                <el-option label="03 : 00" value="3"></el-option>
                                <el-option label="04 : 00" value="4"></el-option>
                                <el-option label="05 : 00" value="5"></el-option>
                                <el-option label="06 : 00" value="6"></el-option>
                                <el-option label="07 : 00" value="7"></el-option>
                                <el-option label="08 : 00" value="8"></el-option>
                                <el-option label="09 : 00" value="9"></el-option>
                                <el-option label="10 : 00" value="10"></el-option>
                                <el-option label="11 : 00" value="11"></el-option>
                                <el-option label="12 : 00" value="12"></el-option>
                                <el-option label="13 : 00" value="13"></el-option>
                                <el-option label="14 : 00" value="14"></el-option>
                                <el-option label="15 : 00" value="15"></el-option>
                                <el-option label="16 : 00" value="16"></el-option>
                                <el-option label="17 : 00" value="17"></el-option>
                                <el-option label="18 : 00" value="18"></el-option>
                                <el-option label="19 : 00" value="19"></el-option>
                                <el-option label="20 : 00" value="20"></el-option>
                                <el-option label="21 : 00" value="21"></el-option>
                                <el-option label="22 : 00" value="22"></el-option>
                                <el-option label="23 : 00" value="23"></el-option>
                                <el-option label="24 : 00" value="24"></el-option>
                            </el-select>
                        </el-col>
                        <el-col :span="2" style="text-align:center;float:right;">
                            <el-col v-text="totalTime"></el-col>
                        </el-col>
                    </el-form-item>

                    <el-form-item label="教员:">
                            <el-col :span="12" style="width: 100%;">
                                <el-autocomplete
                                    class="inline-input"
                                    v-model="form.instructor"
                                    :fetch-suggestions="fetchMostInstructors"
                                    autocomplete="false"
                                    @select="handlerFocusInstructor"
                                    style="width:100%;"
                                ></el-autocomplete>
                            </el-col>
                    </el-form-item>

                    <el-form-item label="学员:">
                            <el-col :span="12" style="width: 100%;">
                                <el-autocomplete
                                    class="inline-input"
                                    v-model="form.students"
                                    :fetch-suggestions="fetchStudentsNames"
                                    @select="handlerFocusInstructor"
                                    style="width:100%;"
                                ></el-autocomplete>
                            </el-col>
                    </el-form-item>

                    <el-form-item label="训练类型:">
                        <el-col :span='12' style="width:100%;">
                            <el-autocomplete
                                v-model="form.mode"
                                :fetch-suggestions='trainingModeSuggestions'
                                @select="handlerFocusInstructor"
                                style="width: 100%;"
                                class="inline-input">

                            </el-autocomplete>
                        </el-col>
                    </el-form-item>

                    <el-form-item label="训练机组号:">
                        <el-input placeholder="" v-model="form.crew"></el-input>
                    </el-form-item>
                    </el-col>

                    <el-col :span="8">
                        <el-form-item label="训练单:">
                            <el-col class="uploadImage-zone">
                                <div @click="uploadTrainingScanImage">
                                    <div v-show="!imgUploaded">
                                        <div>
                                            <svg t="1552455583702" style="width:100%;height:100%;" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="4042" xmlns:xlink="http://www.w3.org/1999/xlink">
                                                <path fill="#d9d9d9" d="M459.861333 0a28.444444 28.444444 0 0 0-28.416 28.16v401.265778H28.302222c-15.644444 0-28.302222 12.8-28.302222 28.416v106.296889c0 15.701333 12.913778 28.416 28.302222 28.416h403.143111v403.143111c0 15.644444 12.8 28.302222 28.444445 28.302222h106.268444c15.701333 0 28.444444-12.913778 28.444445-28.302222V592.554667h401.208889c15.587556 0 28.188444-12.8 28.188444-28.444445v-106.268444a28.444444 28.444444 0 0 0-28.16-28.444445H594.545778V28.216889c0-15.587556-12.8-28.188444-28.416-28.188445h-106.296889z" p-id="4043"></path>
                                            </svg>                                          
                                        </div>
                                    </div> 
                                    <div v-show="imgUploaded">
                                        <img v-bind:src="trainingImg.src" v-bind:alt="trainingImg.alt" max-height="300" width="550">
                                    </div>
                                </div>
                            </el-col>
                            <el-col>
                                <el-button style="margin-top:10px;" size="small" type="primary" @click="uploadTrainingScanImage">点击上传</el-button>
                            </el-col>
                            <el-col>
                                <input type="file" accept="image/jpeg" style="display:none;" id="trainingScanImage" @change="preShowTrainingImage($event)">
                            </el-col>
                        </el-form-item>

                        <el-row style="margin-top:50px;">
                            <el-form-item label="">
                                <el-button type="primary" @click="saveTriningList">保存训练单</el-button>
                            </el-form-item>
                        </el-row>
                    </el-col>

                </el-row>
            </el-form>
        </div>
    </div>
</template>

<script>
import loveuerMenu from "../uMenu.vue";

export default {
    data() {
        return {
            form: {
                date:　"",
                sim: "",
                airline: "",
                session: '',
                ts: "",
                te: "",
                instructor: "",
                students: "",
                mode: "",
                crew: "",
                img: null,
            },
            imgUploaded: false,
            totalTime: 0 + "Hs",
            trainingImg: {
                src: "",
                alt: "",
            }
        };
    },
    components: {
        "loveuer-menu": loveuerMenu,
    },
    methods: {
        showPickedDate: function(e) {
            console.log(e);
        },
        calTotalTime: function() {
            if (this.form.start !== "" && this.form.end !== "") {
                let totalTimeInterval = parseInt(this.form.end) - parseInt(this.form.start);
                if (totalTimeInterval < 0) {
                    totalTimeInterval += 24;
                };
                this.totalTime = totalTimeInterval + "Hs";
            };
        },
        fetchMostInstructors: function(queryString, cb) {
            let mockInstroctors =  [
                {"value":"vue"},
                {"value":"element"},
                {"value":"react"},
                {"value":"angular"},
                {"value":"bootstrap"},
            ];
            cb(mockInstroctors);
        },
        fetchStudentsNames: function(queryString, cb) {
            console.log(queryString); // queryString 为你相当于 keyup 时的 value
            let mockStuendts = [
                {"value":"xxx,yyy"},
                {"value":"aaa,bbb"},
                {"value":"mmm,nnn"},
            ];
            cb(mockStuendts);
        },
        trainingModeSuggestions: function(qs, cb) {
            let tms = [
                {value:'复训'}, {value:'熟练检查'}, {value:'升级训练'}, {value:'升级检查'}, {value:'补充训练'}
            ]
            cb(tms);
        },
        handlerFocusInstructor: function() {
            console.log("select func");
        },
        uploadTrainingScanImage: () => {
            document.querySelector("#trainingScanImage").click();
        },
        preShowTrainingImage: function(e) {
            let imgFile = e.target.files[0];
            let imgReader = new FileReader();
            imgReader.onload = t => {
                this.trainingImg.src = t.target.result;
            };
            imgReader.readAsDataURL(imgFile);
            console.log("uploaded");
            this.imgUploaded = true;
            this.trainingImgalt = this.form.date + this.form.start + ' - ' + this.form.end;
        },
        saveTriningList: function() {
            console.log("save list");
        },
        
    },
    mounted() {
        this.$refs.uMenu.defaultActive = "/works/trainingRecorder/add";
    },
};
</script>

<style scoped>

.uploadImage-zone{
    height: 300px;
    width: 550px;
    border: 1px dashed #d9d9d9;
    border-radius: 6px;
    cursor: pointer;
}
.uploadImage-zone:hover{
    border-color: #409EFF;
}
.uploadImage-zone > div {
    height: 100%;
    width: 100%;
    display: flex;
    align-items: center;
    justify-content: center;
}
.uploadImage-zone > div > div:first-child {
    height: 40px;
    width: 40px;
}
.uploadImage-zone > div > div:last-child {
    height: 100%;
    width: 100%;
    display: flex;
    align-items: center;
    justify-content: center;
}
.uploadImage-zone > div > div:last-child > img {
    width: 540px;
    max-height: 290px;
}
</style>

