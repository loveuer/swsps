<template>
    <div class="container" :style="{display:ifshowStr}" @keydown.esc="quit" tabindex="0">
        <img :src="imgSrc" @click="openImg">
    </div>
</template>

<script>
export default {
    data() {
        return {
            imgSrc: "",
            ifshow: false,
        };
    },
    methods: {
        openImg: function() {

        },
        quit: function() {
            console.log('quit');
            this.ifshow = false;
        },
    },
    computed: {
        ifshowStr: function() {
            if (this.ifshow) {
                return 'block';
            } else {
                return 'none';
            };
        },
    }
};
</script>


<style scoped>
.container{
    width:100%;
    height:100%;
    position: fixed;
    z-index: 999;
    top: 0;
    left: 0;
    background-color: rgba(0,0,0,0.8);
    display: flex;
    justify-content: center;
    text-align: center;
}
.container > img {
    max-width: 90%;
    max-height: 80%;
    margin-top: 2%;
}
</style>
