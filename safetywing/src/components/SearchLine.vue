<template>
    <div id="search">
        <div class="search-div">
            <div class="search-ipt" :style="{ border:'2px solid '+this.$color }">
                <div>
                    <input type="text" @focus="showHistorySearch" @keypress.enter="doSearch">
                </div>
                <span @click="doSearch">
                    <svg t="1542184264203" style="width:30;height:30;position:relative;left:20%;top:20%;" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="5023" xmlns:xlink="http://www.w3.org/1999/xlink">
                        <path v-bind:fill="this.$color" d="M1013.245844 961.315332l-225.864336-225.791193a443.170989 443.170989 0 0 0 105.910479-288.913254 443.756129 443.756129 0 0 0-130.778962-315.829729A443.682987 443.682987 0 0 0 446.610154 0.002194a443.682987 443.682987 0 0 0-315.82973 130.778962A443.756129 443.756129 0 0 0 0.001463 446.684028c0 119.368717 46.445548 231.496316 130.778961 315.829729a443.682987 443.682987 0 0 0 315.82973 130.852104 443.536702 443.536702 0 0 0 288.840111-105.910479l225.791193 225.864336a36.425013 36.425013 0 0 0 26.038765 10.678819 36.717583 36.717583 0 0 0 25.965621-62.683205zM182.78481 710.728799a373.61238 373.61238 0 0 1 0-527.943258 371.125532 371.125532 0 0 1 264.044771-109.421323c99.766501 0 193.389024 38.765576 264.044772 109.421323a373.61238 373.61238 0 0 1 0 527.943258 374.051236 374.051236 0 0 1-528.089543 0z"></path>
                    </svg>
                </span>
            </div>
            <div class="search-drop" v-if="this.lastHistory" :style="dropMenuBorder">
                <template v-for="(his, index) of this.lastHistory">
                    <div v-text="his" v-bind:key="index" class="drop-oneHis" @mouseenter="hoverOneHis($event)" @mouseleave="unHoverOneHis($event)" @click="searchByHis($event)"></div>
                </template>
            </div>
        </div>
    </div>
</template>

<script>
export default {
    data() {
        return {
            lastHistory: [],
            dropMenuBorder: {
                borderRight: "1px solid " + this.$color,
                borderBottom: "1px solid " + this.$color,
                borderLeft: "1px solid " + this.$color,
            },
        }
    },
    methods: {
        showHistorySearch: function() {
            this.$http.get("/api/history/most/" + this.$user.id)
                .then(resp => {
                    this.lastHistory = resp.data;
            });
            for(let ohis of document.querySelectorAll(".drop-oneHis")){
                ohis.style.color = this.$color;
            }
            document.querySelector(".search-drop").style.display = "block";
        },
        doSearch: function() {
            let keyWord = document.querySelector("input").value;
            document.querySelector("input").value = "";
            console.log(keyWord);
            // ******************************************************
            // do search
        },
        hoverOneHis: function(e) {
            e.target.style.backgroundColor = this.$color;
            e.target.style.color = "#fff";
        },
        unHoverOneHis: function(e) {
            e.target.style.backgroundColor = "#fff";
            e.target.style.color = this.$color;
        },
        searchByHis: function(e) {
            document.querySelector("input").value = e.target.innerHTML;
            this.doSearch();
        }
    },
    mounted() {
        document.addEventListener("click", function(event) {
            if(!myContains(event.target, "search-div")) {
                document.querySelector(".search-drop").style.display = "none";
            };
        });
        this.$http.get("/api/history/most/" + this.$user.id)
            .then(resp => {
                this.lastHistory = resp.data;
        });
    },
    
};

function myContains(target, className) {
    while (true){
        if(target == document.body){
            return false;
        };
        if(target.classList.contains(className)){
            return true;
        } else {
            target = target.parentNode;
        };
    };
};
</script>

<style scoped>
#search{
    width: 100%;
    min-width: 700px;
    min-height: 50px;
    height: 70px;
    display: flex;
    align-items: center;
    justify-content: center;
}
.search-div {
    height: 50px;
    width: 700px;
}
.search-ipt{
    height: 100%;
    width: 100%;
    display: inline-flex;
}
.search-ipt > div{
    height: 100%;
    width: 100%;
    display: flex;
    align-items: center;
}
.search-ipt > div > input {
    height: 40px;
    width: 644px;
    outline: none;
    font-size: 1.25em;
    text-indent: 0.5em;
    border: none;
}
.search-ipt > span {
    height: 50px;
    width: 50px;
    cursor: pointer;
}
.search-drop{
    display: none;
    position: relative;
    width: 700px;
    background-color: #fff;
}
.drop-oneHis{
    padding: 5px;
    cursor: pointer;
}
</style>
