<template>
    <div class="quill-editor-example" :style="{width:'100%', height:'100%'}">
        <quill-editor 
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
            contentCode: "",
            editorOption: {
                theme: 'snow',
                placeholder: 'log here',
                rows: '10',
                modules: {
                    toolbar: [
                        ['bold', 'italic', 'underline', 'strike'],
                        [{ 'size': [false, 'large', 'huge'] }],
                        [{ 'list': 'ordered'}, { 'list': 'bullet' }],
                        [{ 'color': [] }, { 'background': [] }],
                        [{ 'align': [] }],
                        ['clean']
                    ],
                },
            },
        };
    },
    components: {
        quillEditor,
    },
    methods: {
        onEditorChange({ editor, html, text }) {
            this.contentCode = html;
        },
    }
};
</script>

<style scoped>
.quill-editor-example {
    min-height: 300px;
}
</style>

