<template>
    <div>
        <div>
            <el-color-picker
                v-model="selColor"
                :predefine="predefineColors">
            </el-color-picker>
        </div>
        <div>
            <el-button @click="aaa">show</el-button>
        </div>
    </div>
</template>

<script>
export default {
    data() {
        return {
            predefineColors: [
                '#ff4500',
                '#ff8c00',
                '#ffd700',
                '#90ee90',
                '#00ced1',
                '#1e90ff',
                '#c71585',
            ],
            selColor: "#000000",
        };
    },
    methods: {
        aaa: function() {
            console.log(this.selColor);
        },
    },
};
</script>
