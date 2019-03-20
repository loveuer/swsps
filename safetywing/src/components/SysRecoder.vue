<template>
    <div>
        <el-timeline>
            <el-timeline-item
                v-for="(m, index) of mock"
                :key="index"
                :timestamp="m.date"
                >
                <span v-text="m.name"></span> : <span v-text="m.age"></span>
            </el-timeline-item>
        </el-timeline>
    </div>
</template>

<script>
export default {
    data() {
        return {
            mock: [
                {date:'2019-03-19', name:'zyp', age:27},
                {date:'2019-03-19', name:'zyp', age:27},
                {date:'2019-03-19', name:'zyp', age:27},
                {date:'2019-03-18', name:'zyp', age:27},
                {date:'2019-03-17', name:'zyp', age:27},
                {date:'2019-03-17', name:'zyp', age:27},
                {date:'2019-03-16', name:'zyp', age:27},
                {date:'2019-03-16', name:'zyp', age:27},
                {date:'2019-03-16', name:'zyp', age:27},
            ]
        }
    },
    computed: {
        
    },
    mounted() {
        
    },
}
</script>

<style scoped>

</style>

