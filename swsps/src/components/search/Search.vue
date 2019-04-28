<template>
    <div id="container">
        <sps-menu :path="'/search'"></sps-menu>
        <div class="search-main">
            <el-card style="width:90%;margin-top:30px;">
                <div slot="header" class="clearfix">
                    <span>
                        <el-autocomplete
                            style="width:400px;"
                            :fetch-suggestions="getSearchSuggestions"
                            placeholder="搜索备件"
                            v-model="search_key">
                            <i slot="prefix" class="el-input__icon el-icon-search"></i>
                        </el-autocomplete>
                    </span>
                    <span style="margin-left:20px;">
                        <el-button type="primary">搜索</el-button>
                    </span>
                </div>
                <div>
                    <el-table>
                        <el-table-column></el-table-column>
                    </el-table>
                </div>
            </el-card>
        </div>
    </div>
</template>

<script>
import SpsMenu from '../Menu.vue';

export default {
    data() {
        return {
            search_key: '',
        };
    },
    methods: {
        getSearchSuggestions(str, cb) {
            let mock = [
                {value:'suggestion 1'},
                {value:'suggestion 2'},
                {value:'suggestion 3'},
            ];
            cb(mock);
            return;
        },
    },
    components: {
        'sps-menu': SpsMenu,
    },
};
</script>

<style scoped>

.search-main {
    width: 100%;
    display: flex;
    justify-content: center;
}
</style>
