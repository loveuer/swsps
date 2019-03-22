<template>
    <div>
        <div>
            <loveuer-menu ref="uMenu"></loveuer-menu>
        </div>

        <div style="margin-top:30px;">
            <el-row style="width:100%;">
                <el-input
                    placeholder="请输入内容"
                    style="width:300px;margin-left:30px;"
                    clearable
                    @keypress.enter.native="doSearch"
                    v-model="searchKey">
                    <i slot="prefix" class="el-input__icon el-icon-search"></i>
                </el-input>
            </el-row>
        </div>
    </div>
</template>

<script>
import loveuerMenu from "../uMenu.vue";

export default {
    data() {
        return {
            searchKey: "",
        };
    },
    methods: {
        doSearch: function() {
            console.log("search: ", this.searchKey);
        },
    },
    mounted() {
        this.$refs.uMenu.defaultActive = "/works/spsRecorder/search";
    },
    components: {
        "loveuer-menu": loveuerMenu,
    },
};
</script>