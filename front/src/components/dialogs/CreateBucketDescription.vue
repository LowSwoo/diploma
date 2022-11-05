<template>
  <v-dialog v-model="showCreateDescriptionDialog" persistent fullscreen>
    <v-card>
      <v-card-title>
        <span class="text-h5">Задать описание расчёта</span>
      </v-card-title>
      <v-card-text>
        <v-container>
          <v-row>
            <v-col cols="12">
              <v-container v-if="editor">
                <v-btn @click="editor.chain().focus().toggleBold().run()"
                  :disabled="!editor.can().chain().focus().toggleBold().run()"
                  :class="{ 'is-active': editor.isActive('bold') }">
                  bold
                </v-btn>
                <v-btn @click="editor.chain().focus().toggleItalic().run()"
                  :disabled="!editor.can().chain().focus().toggleItalic().run()"
                  :class="{ 'is-active': editor.isActive('italic') }">
                  italic
                </v-btn>
                <v-btn @click="editor.chain().focus().toggleStrike().run()"
                  :disabled="!editor.can().chain().focus().toggleStrike().run()"
                  :class="{ 'is-active': editor.isActive('strike') }">
                  strike
                </v-btn>
                <v-btn @click="editor.chain().focus().toggleCode().run()"
                  :disabled="!editor.can().chain().focus().toggleCode().run()"
                  :class="{ 'is-active': editor.isActive('code') }">
                  code
                </v-btn>
                <v-btn @click="editor.chain().focus().unsetAllMarks().run()">
                  clear marks
                </v-btn>
                <v-btn @click="editor.chain().focus().clearNodes().run()">
                  clear nodes
                </v-btn>
                <v-btn @click="editor.chain().focus().setParagraph().run()"
                  :class="{ 'is-active': editor.isActive('paragraph') }">
                  paragraph
                </v-btn>
                <v-btn @click="editor.chain().focus().toggleHeading({ level: 1 }).run()"
                  :class="{ 'is-active': editor.isActive('heading', { level: 1 }) }">
                  h1
                </v-btn>
                <v-btn @click="editor.chain().focus().toggleHeading({ level: 2 }).run()"
                  :class="{ 'is-active': editor.isActive('heading', { level: 2 }) }">
                  h2
                </v-btn>
                <v-btn @click="editor.chain().focus().toggleHeading({ level: 3 }).run()"
                  :class="{ 'is-active': editor.isActive('heading', { level: 3 }) }">
                  h3
                </v-btn>
                <v-btn @click="editor.chain().focus().toggleHeading({ level: 4 }).run()"
                  :class="{ 'is-active': editor.isActive('heading', { level: 4 }) }">
                  h4
                </v-btn>
                <v-btn @click="editor.chain().focus().toggleHeading({ level: 5 }).run()"
                  :class="{ 'is-active': editor.isActive('heading', { level: 5 }) }">
                  h5
                </v-btn>
                <v-btn @click="editor.chain().focus().toggleHeading({ level: 6 }).run()"
                  :class="{ 'is-active': editor.isActive('heading', { level: 6 }) }">
                  h6
                </v-btn>
                <v-btn @click="editor.chain().focus().toggleBulletList().run()"
                  :class="{ 'is-active': editor.isActive('bulletList') }">
                  bullet list
                </v-btn>
                <v-btn @click="editor.chain().focus().toggleOrderedList().run()"
                  :class="{ 'is-active': editor.isActive('orderedList') }">
                  ordered list
                </v-btn>
                <v-btn @click="editor.chain().focus().toggleCodeBlock().run()"
                  :class="{ 'is-active': editor.isActive('codeBlock') }">
                  code block
                </v-btn>
                <v-btn @click="editor.chain().focus().toggleBlockquote().run()"
                  :class="{ 'is-active': editor.isActive('blockquote') }">
                  blockquote
                </v-btn>
                <v-btn @click="editor.chain().focus().setHorizontalRule().run()">
                  horizontal rule
                </v-btn>
                <v-btn @click="editor.chain().focus().setHardBreak().run()">
                  hard break
                </v-btn>
                <v-btn @click="editor.chain().focus().undo().run()"
                  :disabled="!editor.can().chain().focus().undo().run()">
                  undo
                </v-btn>
                <v-btn @click="editor.chain().focus().redo().run()"
                  :disabled="!editor.can().chain().focus().redo().run()">
                  redo
                </v-btn>
              </v-container>
              <v-divider></v-divider>
              <editor-content :editor="editor" />
            </v-col>
          </v-row>
        </v-container>
      </v-card-text>
      <v-card-actions>
        <v-spacer></v-spacer>
        <v-btn color="blue darken-1" text @click="CloseDialog">
          Закрыть
        </v-btn>
        <v-btn color="blue darken-1" text @click="Save">
          Сохранить
        </v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>

<script>
import { Editor, EditorContent, getDebugJSON } from '@tiptap/vue-2'
import StarterKit from '@tiptap/starter-kit'
import settings from '@/settings'

export default {
  props: ['showCreateDescriptionDialog', 'bucketName'],
  components: {
    EditorContent,
  },

  data() {
    return {
      editor: null,
    }
  },
  methods: {
    CloseDialog() {
      this.$emit("CloseDescriptionDialog")
    },
    Save() {
      console.log(this.editor.getJSON())
      this.$http.post(settings.ebaHOST + "api/bucket/setDescription", {
        description: this.editor.getJSON(),
        bucketName: this.bucketName
      })
      this.$emit("CloseDescriptionDialog")
    }
  },
  mounted() {
    this.editor = new Editor({
      content: '<p>I’m running Tiptap with Vue.js. 🎉</p>',
      extensions: [
        StarterKit,
      ],
      autofocus: true
    })
  },

  beforeDestroy() {
    this.editor.destroy()
  },
}
</script>