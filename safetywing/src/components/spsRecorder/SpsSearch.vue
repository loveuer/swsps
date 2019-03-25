<template>
    <div>
        <div>
            <loveuer-menu ref="uMenu"></loveuer-menu>
        </div>

        <div style="margin-top:30px;">
            <el-row style="width:100%;">
                <!-- @keypress.enter.native="doSearch" -->
                <span style="margin-left:30px;">
                    <el-button circle type="success" icon="el-icon-plus" @click="ifshowAdd=true"></el-button>
                </span>
                <span>
                    <el-input
                        placeholder="请输入内容"
                        style="width:300px;margin-left:10px;"
                        @keypress.enter.native="doSearch"
                        clearable
                        v-model="searchKey">
                        <i slot="prefix" class="el-input__icon el-icon-search"></i>
                    </el-input>
                </span>
                <span v-show="sps" style="line-height:40px;margin-left:30px;">
                    <font>共搜索到: </font><font style="color:#F06560" v-text="searchedAmount"></font><font> 个结果</font>
                </span>
            </el-row>
            <el-row class="result-filter">
                <el-checkbox-group v-model="simsList" :min="1">
                    <el-checkbox label="5978"></el-checkbox>
                    <el-checkbox label="5989"></el-checkbox>
                    <el-checkbox label="5008"></el-checkbox>
                    <el-checkbox label="5015"></el-checkbox>
                </el-checkbox-group>
            </el-row>
        </div>
        <!-- content -->
        <div style="width:100%;min-width:1100px;margin-left:30px;">
            <el-table
                :data="filteredSps"
                height="700"
                @row-click="meetSpDetail"
                style="width:100%;">
                <el-table-column type="expand">
                    <template slot-scope="props">
                        <el-form label-position="left" class="demo-table-expand" label-width="120px">
                            <el-form-item label="原始模拟机"><span>{{ props.row.orgsim }}</span></el-form-item>
                            <el-form-item label="备注"><span>{{ props.row.comment }}</span></el-form-item>
                            <el-form-item label="标签"><span>{{ props.row.tags }}</span></el-form-item>
                        </el-form>
                    </template>
                </el-table-column>
                <el-table-column label="名称" prop="name" width="400"></el-table-column>
                <el-table-column label="P/N" prop="pn" width="350"></el-table-column>
                <el-table-column label="S/N" prop="sn" width="250"></el-table-column>
                <el-table-column label="状态" prop="status" width="100"></el-table-column>
                <el-table-column label="模拟机" prop="nowsim" width="100"></el-table-column>
                <el-table-column label="位置" prop="pos" width="150"></el-table-column>
            </el-table>
        </div>
        <div>
            <el-dialog title="新增备件" :visible.sync="ifshowAdd" width="935px" style="margin-top:-50px;">
                <el-form label-position="left" label-width="70px">
                    <el-row>
                        <el-col :span="12">
                            <el-form-item label="名称" style="width:350px;">
                                <el-input v-model="newsp.name"></el-input>
                            </el-form-item>
                        </el-col>
                        <el-col :span="12" style="margin-left:-50px;">
                            <el-form-item label="P/N" style="width:350px;">
                                <el-input v-model="newsp.pn"></el-input>
                            </el-form-item>
                        </el-col>
                    </el-row>
                    <el-row>
                        <el-col :span="12">
                            <el-form-item label="S/N" style="width:350px;">
                                <el-input v-model="newsp.sn"></el-input>
                            </el-form-item>
                        </el-col>
                        <el-col :span="12" style="margin-left:-50px;">
                            <el-form-item label="位置" style="width:350px;">
                                <el-input v-model="newsp.pos"></el-input>
                            </el-form-item>
                        </el-col>
                    </el-row>
                    <el-row>
                        <el-col :span="12">
                            <el-form-item label="原模拟机" style="width:350px;">
                                <el-select v-model="newsp.orgsim" style="width:100%;">
                                    <el-option label="5978" value="5978"></el-option>
                                    <el-option label="5989" value="5989"></el-option>
                                    <el-option label="5008" value="5008"></el-option>
                                    <el-option label="5015" value="5015"></el-option>
                                </el-select>
                            </el-form-item>
                        </el-col>
                        <el-col :span="12" style="margin-left:-50px;">
                            <el-form-item label="现模拟机" style="width:350px;">
                                <el-select v-model="newsp.nowsim" style="width:100%;">
                                    <el-option label="5978" value="5978"></el-option>
                                    <el-option label="5989" value="5989"></el-option>
                                    <el-option label="5008" value="5008"></el-option>
                                    <el-option label="5015" value="5015"></el-option>
                                </el-select>
                            </el-form-item>
                        </el-col>
                    </el-row>
                    <el-row>
                        <el-col :span="12">
                            <el-form-item label="状态" style="width:350px;">
                                <el-select v-model="newsp.status" style="width:100%;">
                                    <el-option label="备件" value="备件"></el-option>
                                    <el-option label="使用" value="使用"></el-option>
                                    <el-option label="故障" value="故障"></el-option>
                                    <el-option label="维修" value="维修"></el-option>
                                    <el-option label="废弃" value="废弃"></el-option>
                                </el-select>
                            </el-form-item>
                        </el-col>
                        <el-col :span="12" style="margin-left:-50px;">
                            <el-form-item label="数量" style="width:350px;">
                                <el-input-number v-model="newsp.num" :min="0" style="width:100%;"></el-input-number>
                            </el-form-item>
                        </el-col>
                    </el-row>
                    <el-row>
                        <el-form-item label="是否耗材" style="width:350px;">
                            <el-radio v-model="newsp.is_consumable" label="0">不是耗材</el-radio>
                            <el-radio v-model="newsp.is_consumable" label="1">是耗材</el-radio>
                        </el-form-item>
                    </el-row>
                    <el-row>
                        <el-form-item label="备注" style="width:750px;">
                            <el-input v-model="newsp.comment"></el-input>
                        </el-form-item>
                    </el-row>
                    <el-row>
                        <el-col :span="12">
                            <el-form-item label="图片1" style="width:350px;">
                                <div class="newsp-img" @click="clickImgInput('#newspimg1-input')">
                                    <!-- 没有选择图片的时候 -->
                                    <div v-show="!newsp.img1" class="newspimg-none">
                                        <div>
                                            <svg t="1552455583702" style="width:100%;height:100%;" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="4042" xmlns:xlink="http://www.w3.org/1999/xlink">
                                                <path fill="#d9d9d9" d="M459.861333 0a28.444444 28.444444 0 0 0-28.416 28.16v401.265778H28.302222c-15.644444 0-28.302222 12.8-28.302222 28.416v106.296889c0 15.701333 12.913778 28.416 28.302222 28.416h403.143111v403.143111c0 15.644444 12.8 28.302222 28.444445 28.302222h106.268444c15.701333 0 28.444444-12.913778 28.444445-28.302222V592.554667h401.208889c15.587556 0 28.188444-12.8 28.188444-28.444445v-106.268444a28.444444 28.444444 0 0 0-28.16-28.444445H594.545778V28.216889c0-15.587556-12.8-28.188444-28.416-28.188445h-106.296889z" p-id="4043"></path>
                                            </svg>
                                        </div>
                                    </div>
                                    <!-- 选择了图片, 进行预览 -->
                                    <div v-show="newsp.img1" class="newspimg-got">
                                        <img :src="imgPreview1Src" alt="图片1" max-width="280" max-height="180">
                                    </div>
                                    <!-- 隐藏的 input 图片区域 -->
                                    <div>
                                        <input type="file" accept="image/jpeg" style="display:none;" id="newspimg1-input" @change="preViewImg(1, $event)">
                                    </div>
                                </div>
                            </el-form-item>
                        </el-col>
                        <el-col :span="12" style="margin-left:-50px;">
                            <el-form-item label="图片2" style="width:350px;">
                                <div class="newsp-img" @click="clickImgInput('#newspimg2-input')">
                                    <!-- 没有选择图片的时候 -->
                                    <div v-show="!newsp.img2" class="newspimg-none">
                                        <div>
                                            <svg t="1552455583702" style="width:100%;height:100%;" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="4042" xmlns:xlink="http://www.w3.org/1999/xlink">
                                                <path fill="#d9d9d9" d="M459.861333 0a28.444444 28.444444 0 0 0-28.416 28.16v401.265778H28.302222c-15.644444 0-28.302222 12.8-28.302222 28.416v106.296889c0 15.701333 12.913778 28.416 28.302222 28.416h403.143111v403.143111c0 15.644444 12.8 28.302222 28.444445 28.302222h106.268444c15.701333 0 28.444444-12.913778 28.444445-28.302222V592.554667h401.208889c15.587556 0 28.188444-12.8 28.188444-28.444445v-106.268444a28.444444 28.444444 0 0 0-28.16-28.444445H594.545778V28.216889c0-15.587556-12.8-28.188444-28.416-28.188445h-106.296889z" p-id="4043"></path>
                                            </svg>
                                        </div>
                                    </div>
                                    <!-- 选择了图片, 进行预览 -->
                                    <div v-show="newsp.img2" class="newspimg-got">
                                        <img :src="imgPreview2Src" alt="图片2" max-width="280" max-height="180">
                                    </div>
                                    <!-- 隐藏的 input 图片区域 -->
                                    <div>
                                        <input type="file" accept="image/jpeg" style="display:none;" id="newspimg2-input" @change="preViewImg(2, $event)">
                                    </div>
                                </div>
                            </el-form-item>
                        </el-col>
                    </el-row>
                    
                </el-form>
                <el-row>
                    <el-button type="primary" @click="postNewSp">确 定</el-button>
                    <el-button @click="ifshowAdd = false">取 消</el-button>
                </el-row>
            </el-dialog>
        </div>
        <div style="width:100%;height:120px;zoom:100%"></div>
    </div>
</template>

<script>
import loveuerMenu from "../uMenu.vue";

export default {
    data() {
        return {
            searchKey: "", 
            simsList: ['5978', '5989', '5008', '5015'],
            sps: [],
            ifshowAdd: false,
            newsp: {
                name: '', pn: '', sn: '', pos: '', orgsim: '', nowsim: '', status:'', num: 1, is_consumable: '0', comment: '', img1: false, img2: false,
            },
            imgPreview1Src: '',
            imgPreview2Src: '',
        };
    },
    methods: {
        doSearch: function() {
            this.$router.push("/works/spsRecorder/search/" + this.searchKey);
            console.log("search: ", this.searchKey);
            this.$http.get("/api/sps/search/" + this.searchKey)
                .then(resp => {
                    this.sps = resp.data;
                })
                .catch(error => { console.log("sps search http error: ", error); });
        },
        meetSpDetail: function(row) {
            this.$router.push("/works/spsRecorder/detail/" + row.id);
        },
        clickImgInput: function(imgid) {
            document.querySelector(imgid).click();
        },
        preViewImg: function(which, e) {
            let imgFile = e.target.files[0];
            let imgReader = new FileReader();

            imgReader.onload = t => {
                if (which === 1) {
                    this.imgPreview1Src = t.target.result;
                } else if (which === 2) {
                    this.imgPreview2Src = t.target.result;
                };
            };
            imgReader.readAsDataURL(imgFile);
            if (which === 1) {
                this.newsp.img1 = true;
            } else if (which === 2) {
                this.newsp.img2 = true;
            };
        },
        postNewSp: function() {
            console.log(this.newsp);
            this.ifshowAdd = false;

            // mock success
            this.$confirm('上传成功, 是否继续上传?', '提示', {
                confirmButtonText: '继续',
                cancelButtonText: '取消',
                type: 'info',
            }).then(() => {
                console.log("yes");
                this.ifshowAdd = true;
                this.newsp.sn = '';
                // this.newsp.img1 = '';
                // this.newsp.img2 = '';
            }).catch(() => {
                console.log('no');
            });
        },
    },
    computed: {
        filteredSps: function() {
            if (this.sps) {
                return this.sps.filter(sp => this.simsList.includes(sp.nowsim));
            } else {
                return [];
            };
        },
        searchedAmount: function() {
            if (this.sps) {
                return this.sps.length;
            } else {
                return 0;
            };
        },
    },
    mounted() {
        this.$refs.uMenu.defaultActive = "/works/spsRecorder/search";
        if (this.$route.params.skey) {
            this.searchKey = this.$route.params.skey;
            this.doSearch();
        };
    },
    components: {
        "loveuer-menu": loveuerMenu,
    },
};
</script>

<style scoped>
.result-filter{
    margin-left: 32px;
    margin-top: 20px;
}
.demo-table-expand {
    font-size: 0;
}
.demo-table-expand label {
    width: 90px;
    color: #99a9bf;
}
.demo-table-expand .el-form-item {
    margin-right: 0;
    margin-bottom: 0;
    width: 50%;
}
.newsp-img {
    height: 180px;
    width:280px;
    border: 1px dashed #ddd;
    border-radius: 10px;
    cursor: pointer;
}
.newsp-img:hover {
    border: 1px dashed #409EFF;
}
.newspimg-none {
    height: 100%;
    width: 100%;
    display: flex;
    align-items: center;
    justify-content: center;
}
.newspimg-none > div {
    height: 50px;
    width: 50px;
}
.newspimg-got {
    border-radius: 10px;
    display: flex;
    align-items: center;
    justify-content: center;
}
.newspimg-got > img {
    max-width: 280px;
    max-height: 180px;
    border-radius: 10px;
    height: 100%;
    width: auto;
}
</style>
