<template>
  
    <div>
        <!-- quill-editor -->
        <quill-editor v-model="content"
            :options="editorOption"
            style="width:100%;height:250px;zoom:100%;margin-bottom:50px;"
            v-on:change="quillContentChg"
            ref="myQuillEditor">

            <div id="toolbar" slot="toolbar">
                <button class="ql-bold" title="粗体"></button>
                <button class="ql-italic" title="斜体"></button>
                <button class="ql-underline" title="下划线"></button>
                <select class="ql-size" title="字体大小">
                    <option value="small"></option>
                    <option selected></option>
                    <option value="large"></option>
                    <option value="huge"></option>
                </select>
                <button class="ql-list" value="ordered" title="列表"></button>
                <select class="ql-color" title="字体颜色"></select>
                <select class="ql-background" title="背景颜色"></select>
                <select class="ql-align" title="对齐">
                    <option selected></option>
                    <option value="center"></option>
                    <option value="right"></option>
                    <option value="justify"></option>
                </select>
            <!-- You can also add your own -->
                <button @click="customButtonClick" style="width:80px;">插入备件</button>
            </div>
        </quill-editor>
    </div>
</template>

<script>
import { quillEditor } from './Quill/index.js';
import 'quill/dist/quill.core.css';
import 'quill/dist/quill.snow.css';

import _Quill from 'quill';

export default {
    data() {
        return {
            content: "",
            editorOption: {
                modules: {
                    toolbar: '#toolbar'
                },
                theme: 'snow',
                placeholder: 'log here',
                rows: '10',
            },
            textIndex: 0,
        };
    },
    components: {
        "quill-editor": quillEditor,
    },
    methods: {
        updateIndex: function() {
            this.textIndex = this.$refs.myQuillEditor.quill.getLength();
        },
        customButtonClick() {
            this.$emit("findpn", '');
        },
        quillContentChg({editor, html, text}) {
            if (text.includes('pn:')) {
                this.$emit('findpn', 'pn:');
            } else if (text.includes('PN:')) {
                this.$emit('findpn', 'PN:');
            };
        },
        setCursorEnd: function() {
            this.updateIndex();
            this.$refs.myQuillEditor.quill.setSelection(this.textIndex-1);
            this.updateIndex();
        },
        insertSps: function(row, findkey) {
            let val = {href: `/works/spsRecorder/detail/${row.id}`, pn: row.pn, sn: row.sn,};
            if (findkey !== "") {
                console.log('findkey 不等于 空');
                let regex = /pn:/gi;
                let text = this.content;
                this.content = text.replace(regex, '');
            };
            // 没办法,,,不设置一个延时, 老是替换不掉pn:
            setTimeout(() => {
                this.setCursorEnd();
                this.$refs.myQuillEditor.quill.insertEmbed(this.textIndex-1, 'link', val);
                this.setCursorEnd();
            }, 300);
        },
    },
    mounted() {
        
    },
};
</script>

<style scoped>

</style>