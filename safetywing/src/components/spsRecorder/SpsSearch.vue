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
            <el-dialog title="新增备件" :visible.sync="ifshowAdd">
                <el-form label-position="left" label-width="100px">
                    <el-form-item label="名称" style="width:400px;">
                        <el-input v-model="newsp.name"></el-input>
                    </el-form-item>
                    <el-form-item label="P/N" style="width:400px;">
                        <el-input v-model="newsp.pn"></el-input>
                    </el-form-item>
                    <el-form-item label="S/N" style="width:400px;">
                        <el-input v-model="newsp.sn"></el-input>
                    </el-form-item>
                    <el-form-item label="原模拟机" style="width:400px;">
                        <el-select v-model="newsp.orgsim" style="width:100%;">
                            <el-option label="5978" value="5978"></el-option>
                            <el-option label="5989" value="5989"></el-option>
                            <el-option label="5008" value="5008"></el-option>
                            <el-option label="5015" value="5015"></el-option>
                        </el-select>
                    </el-form-item>
                    <el-form-item label="现模拟机" style="width:400px;">
                        <el-select v-model="newsp.nowsim" style="width:100%;">
                            <el-option label="5978" value="5978"></el-option>
                            <el-option label="5989" value="5989"></el-option>
                            <el-option label="5008" value="5008"></el-option>
                            <el-option label="5015" value="5015"></el-option>
                        </el-select>
                    </el-form-item>
                    <el-form-item label="是否耗材" style="width:400px;">
                        <el-radio v-model="newsp.is_consumable" label="0">不是耗材</el-radio>
                        <el-radio v-model="newsp.is_consumable" label="1">是耗材</el-radio>
                    </el-form-item>
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
                name: '', pn: '', sn: '', orgsim: '', nowsim: '', is_consumable: '0',
            },
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
</style>
