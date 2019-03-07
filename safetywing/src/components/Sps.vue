<template>
    <div>
        <my-nav />
        <div id="sps-container">
            <div class="split-line" :style="{ backgroundColor:this.$store.state.customColor }"></div>
            <div class="surface">
                <div>
                    <h1>占位</h1>
                </div>
            </div>
            <div class="search">
                <mySearch />
            </div>
            <div class="more">
                <div>
                    <h1>占位2</h1>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
import myNav from "./MyNav.vue"
import mySearch from "./SearchLine"
export default {
    data() {
        return {
            mostHistory: [],
        }
    },
    components: {
        myNav,
        mySearch,
    },
    mounted() {
        this.$http.get("/api/history/search/last/" + this.$store.state.user.id)
            .then(resp => {
                this.mostHistory = JSON.parse(resp.data['val']);
            });
        //
    },
    methods: {
        showHistorySearch: function() {
            document.querySelector(".search-drop").style.display = "block";
        },
    },
}
</script>

<style scoped>
#sps-container{
    width: 100%;
    min-height: 100%;
 }
.split-line{
    width: 100%;
    height: 10px;
    margin-top: 5px;
}
.surface{
    height: 200px;
    width: 100%;
    background-color: lightgray;
    display: flex;
    align-items: center;
    justify-content: center;
}
.search{
    height: 100px;
    width: 100%;
    display: flex;
    align-items: center;
    justify-content: center;
}
.more{
    height: 400px;
    width: 100%;
    background-color: lightseagreen;
    display: flex;
    align-items: center;
    justify-content: center;
}
@media screen and (min-width: 500px){
    
}
</style>
