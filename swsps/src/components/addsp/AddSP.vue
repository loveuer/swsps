<template>
    <div id="container">
        <sps-menu :path="'/addsp'"></sps-menu>
        <div style="width:100%;height:30px;zoom:100%;"></div>
        <div class="main-container">
            <div style="width:65%;min-width:700px;">
                <el-card class="box-card">
                    <div slot="header" class="clearfix">
                        <span>新增备件</span>
                        <el-tooltip content="重置表单" placement="top">
                            <el-button icon="el-icon-refresh" circle style="float:right;margin-top:-8px;"></el-button>
                        </el-tooltip>
                    </div>
                    <div>
                        <el-form ref="form" :model="newsp" label-width="120px" style="width:100%;">
                            <el-form-item label="名称">
                                <el-autocomplete
                                    style="width:400px;"
                                    class="inline-input"
                                    v-model="newsp.name"
                                    :fetch-suggestions="nameSuggestion"
                                >
                                </el-autocomplete>
                            </el-form-item>
                            <el-form-item label="P/N">
                                <el-autocomplete
                                    style="width:400px;"
                                    class="inline-input"
                                    v-model="newsp.pn"
                                    :fetch-suggestions="pnSuggestion"
                                >
                                </el-autocomplete>
                            </el-form-item>
                            <el-form-item label="S/N">
                                <el-input style="width:400px;"></el-input>
                            </el-form-item>
                            <el-form-item label="原模拟机">
                                <el-select v-model="newsp.orgsim" style="width:400px;">
                                    <el-option v-for="item of sims" :key="item.index" :value="item.sim"></el-option>
                                </el-select>
                            </el-form-item>
                            <el-form-item label="模拟机">
                                <el-select v-model="newsp.sim" style="width:400px;">
                                    <el-option v-for="item of sims" :key="item.index" :value="item.sim"></el-option>
                                </el-select>
                            </el-form-item>
                            <el-form-item label="位置">
                                <el-input v-model="newsp.pos" style="width:400px;"></el-input>
                            </el-form-item>
                            <el-form-item label="数量">
                                <el-input-number v-model="newsp.num" :min="0"></el-input-number>
                            </el-form-item>
                            <el-form-item label="图片(最多两张)">
                                <el-button type="primary" size="small">点击上传</el-button>
                                <div class="upload-img">
                                    <div class="place-img">
                                        <img src="~@/assets/imgs/test.jpg" alt="">
                                    </div>
                                    <div class="place-describe">
                                        <div>IMG 1</div>
                                    </div>
                                    <label class="place-corner">
                                        <div class="upload-success">
                                            <i class="el-icon-check corner-check"></i>
                                        </div>
                                        <div class="delete-upload"  @click="deleteUpload('img1')">
                                            <i class="el-icon-close corner-close"></i>
                                        </div>
                                    </label>
                                </div>
                                <div class="upload-img">
                                    <div class="place-img">
                                        <img src="~@/assets/imgs/test.jpg" alt="">
                                    </div>
                                    <div class="place-describe">
                                        <div>IMG 2</div>
                                    </div>
                                    <label class="place-corner">
                                        <div class="upload-success">
                                            <i class="el-icon-check corner-check"></i>
                                        </div>
                                        <div class="delete-upload" @click="deleteUpload('img2')">
                                            <i class="el-icon-close corner-close"></i>
                                        </div>
                                    </label>
                                </div>
                            </el-form-item>
                            <el-form-item>
                                <el-button type="success">新增备件</el-button>
                            </el-form-item>
                        </el-form>
                    </div>
                </el-card>
            </div>
        </div>
        <div style="width:100%;height:100px;zoom:100%;"></div>
    </div>
</template>

<script>
import spsMenu from "../Menu.vue";

export default {
    data() {
        return {
            newsp: { name: '', pn: '', sn: '', orgsim: '', sim: '', pos: '', num: 0, img1: null, img2: null, },
            sims: [{index:1, sim:'5978'}, {index:2, sim:'5989'}, {index:3, sim:'5008'}, {index:4, sim:'5015'}],
        };
    },
    methods: {
        nameSuggestion: function(inputedStr, cb) {
            let mock_nameSuggestions = [{value:'name_123'},{value:'name_234'},{value:'name_345'},{value:'name_456'}];
            cb(mock_nameSuggestions);
        },
        pnSuggestion: function(inputedStr, cb) {
            let mock_pnSuggestions = [{value:'pn_123'},{value:'pn_234'},{value:'pn_345'},{value:'pn_456'}];
            cb(mock_pnSuggestions);
        },
        deleteUpload: function(which) {
            console.log("delete upload ", which);
        },
    },
    computed: {
        
    },
    components: {
        'sps-menu': spsMenu,
    }
};
</script>

<style scoped>
.main-container{
    width: 100%;
    display: flex;
    justify-content: center;
}
.upload-img {
    border: 1px solid rgb(220,223,230);
    border-radius: 5px;
    padding: 10px;
    width: 380px;
    height: 100px;
    margin-top: 10px;
    position: relative;
    overflow: hidden;
    display: flex;
}
.upload-img:hover {
    border: 1px solid rgb(192,196,204);
}
.upload-img:hover .upload-success {
    display: none;
}
.upload-img:hover .place-corner {
    background: #fff;
    box-shadow: none;
}
.upload-img:hover .delete-upload {
    display: block;
}
.place-img {
    height: 100%;
    width: 150px;
}
.place-img > img {
    width: 150px;
    height: 100%;
}
.place-describe {
    width: 100px;
    height: 100%;
    display: flex;
    justify-content: center;
    align-items: center;
}
.place-corner {
    position: absolute;
    top: -7px;
    right: -17px;
    width: 46px;
    height: 26px;
    background: #13ce66;
    transform: rotate(45deg);
    text-align: center;
    box-shadow: 0 1px 1px #ccc;
    display: block;
}
.upload-success{
    margin-top: -2px;
}
.corner-check {
    transform: rotate(-45deg);
    color: #fff;
    margin-top: 12px;
    font-size: 12px;
}
.delete-upload {
    display: none;
}
.corner-close {
    transform: rotate(-45deg);
    margin-top: 12px;
    font-size: 12px;
    cursor: pointer;
}
</style>
