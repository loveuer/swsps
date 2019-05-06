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
                        <el-form ref="form" :model="newsp" label-width="100px" style="width:100%;">
                            <el-form-item label="名称">
                                <el-autocomplete
                                    style="width:300px;"
                                    class="inline-input"
                                    v-model="newsp.name"
                                    :fetch-suggestions="nameSuggestion"
                                >
                                </el-autocomplete>
                            </el-form-item>
                            <el-form-item label="P/N">
                                <el-autocomplete
                                    style="width:300px;"
                                    class="inline-input"
                                    v-model="newsp.pn"
                                    :fetch-suggestions="pnSuggestion"
                                >
                                </el-autocomplete>
                            </el-form-item>
                            <el-form-item label="原模拟机">
                                <el-select v-model="newsp.orgsim" style="width:300px;">
                                    <el-option v-for="item of sims" :key="item.index" :value="item.sim"></el-option>
                                </el-select>
                            </el-form-item>
                            <el-form-item label="模拟机">
                                <el-select v-model="newsp.sim" style="width:300px;">
                                    <el-option v-for="item of sims" :key="item.index" :value="item.sim"></el-option>
                                </el-select>
                            </el-form-item>
                            <el-form-item label="位置">
                                <el-input v-model="newsp.pos" style="width:300px;"></el-input>
                            </el-form-item>
                            <el-form-item label="数量">
                                <el-input-number v-model="newsp.num" :min="0"></el-input-number>
                            </el-form-item>
                            <div v-for="sn of sns" :key="sn.index">
                                <el-form-item label="S/N">
                                    <el-input></el-input>
                                </el-form-item>
                            </div>
                        </el-form>
                    </div>
                </el-card>
            </div>
        </div>
    </div>
</template>

<script>
import spsMenu from "../Menu.vue";
 import AddSNIMG from "./Add-SN_n_IMG.vue";

export default {
    data() {
        return {
            newsp: { name: '', pn: '', sn: '', orgsim: '', sim: '', pos: '', num: 0, },
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
    },
    computed: {
        sns: function() {
            let list = [];
            for (let i=0;i<this.newsp.sum;i++) {
                list.push({index:i, item:'sn'+i});
            };
            return list
        },
    },
    components: {
        'sps-menu': spsMenu,
        'sn-img': AddSNIMG,
    }
};
</script>

<style scoped>
.main-container{
    width: 100%;
    display: flex;
    justify-content: center;
}
</style>
