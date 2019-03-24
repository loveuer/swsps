<template>
    <div class="quill-editor-example" :style="{width:'100%', height:'100%'}">
        <quill-editor 
            v-model="content"
            ref="myQuillEditor"
            :options="editorOption"
            placeholder="'log here'"
            :style="{height:'250px'}"
            @change="onEditorChange($event)">
        </quill-editor>
    </div>
</template>

<script>
import 'quill/dist/quill.core.css';
import 'quill/dist/quill.snow.css';
import { quillEditor } from 'vue-quill-editor';

export default {
    data() {
        return {
            editorOption: {
                theme: 'snow',
                placeholder: 'log here',
                rows: '10',
                modules: {
                    toolbar: [
                        ['bold', 'italic', 'underline', 'strike'],
                        [{ 'size': [false, 'large', 'huge'] }],
                        [{ 'list': 'ordered'}],
                        [{ 'color': [] }, { 'background': [] }],
                        [{ 'align': [] }],
                    ],
                },
            },
            content: '',
        };
    },
    components: {
        quillEditor,
    },
    methods: {
        onEditorChange({ editor, html, text }) {
            this.content = html;
            if (this.content.includes('pn') || this.content.includes('PN')) {
                this.$emit('findpn', {content: text});
            };
        },
        setCursorEnd: function() {
            console.log("set c");
            this.$refs.myQuillEditor.quill.setSelection(1);
        },
    }
};
</script>

<style scoped>
.quill-editor-example {
    min-height: 300px;
}
</style>

