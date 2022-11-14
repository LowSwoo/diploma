<template>
  <v-dialog v-model="showCreateDescriptionDialog"
  transition="dialog-bottom-transition"

  persistent max-width="90%">
    <v-card min-height="80%">
      <v-card-title >
        <h1 class="text-h5">Редактирование описания расчёта: "{{bucketName}}"</h1>
      </v-card-title>
      <v-card-text>
        <!-- <v-container> -->
          <v-row>
            <v-col cols="12">
              <v-container v-if="editor">
                <v-btn outlined small @click="editor.chain().focus().toggleBold().run()"
                  :disabled="!editor.can().chain().focus().toggleBold().run()"
                  :class="{ 'is-active': editor.isActive('bold') }">
                  bold
                </v-btn>
                <v-btn outlined small @click="editor.chain().focus().toggleItalic().run()"
                  :disabled="!editor.can().chain().focus().toggleItalic().run()"
                  :class="{ 'is-active': editor.isActive('italic') }">
                  italic
                </v-btn>
                <v-btn outlined small @click="editor.chain().focus().toggleStrike().run()"
                  :disabled="!editor.can().chain().focus().toggleStrike().run()"
                  :class="{ 'is-active': editor.isActive('strike') }">
                  strike
                </v-btn>
                <v-btn outlined small @click="editor.chain().focus().toggleCode().run()"
                  :disabled="!editor.can().chain().focus().toggleCode().run()"
                  :class="{ 'is-active': editor.isActive('code') }">
                  code
                </v-btn>
                <v-btn outlined small @click="editor.chain().focus().unsetAllMarks().run()">
                  clear marks
                </v-btn>
                <v-btn outlined small @click="editor.chain().focus().clearNodes().run()">
                  clear nodes
                </v-btn>
                <v-btn outlined small @click="editor.chain().focus().setParagraph().run()"
                  :class="{ 'is-active': editor.isActive('paragraph') }">
                  paragraph
                </v-btn>
                <v-btn outlined small @click="editor.chain().focus().toggleHeading({ level: 1 }).run()"
                  :class="{ 'is-active': editor.isActive('heading', { level: 1 }) }">
                  h1
                </v-btn>
                <v-btn outlined small @click="editor.chain().focus().toggleHeading({ level: 2 }).run()"
                  :class="{ 'is-active': editor.isActive('heading', { level: 2 }) }">
                  h2
                </v-btn>
                <v-btn outlined small @click="editor.chain().focus().toggleHeading({ level: 3 }).run()"
                  :class="{ 'is-active': editor.isActive('heading', { level: 3 }) }">
                  h3
                </v-btn>
                <v-btn outlined small @click="editor.chain().focus().toggleHeading({ level: 4 }).run()"
                  :class="{ 'is-active': editor.isActive('heading', { level: 4 }) }">
                  h4
                </v-btn>
                <v-btn outlined small @click="editor.chain().focus().toggleHeading({ level: 5 }).run()"
                  :class="{ 'is-active': editor.isActive('heading', { level: 5 }) }">
                  h5
                </v-btn>
                <v-btn outlined small @click="editor.chain().focus().toggleHeading({ level: 6 }).run()"
                  :class="{ 'is-active': editor.isActive('heading', { level: 6 }) }">
                  h6
                </v-btn>
                <v-btn outlined small @click="editor.chain().focus().toggleBulletList().run()"
                  :class="{ 'is-active': editor.isActive('bulletList') }">
                  bullet list
                </v-btn>
                <v-btn outlined small @click="editor.chain().focus().toggleOrderedList().run()"
                  :class="{ 'is-active': editor.isActive('orderedList') }">
                  ordered list
                </v-btn>
                <v-btn outlined small @click="editor.chain().focus().toggleCodeBlock().run()"
                  :class="{ 'is-active': editor.isActive('codeBlock') }">
                  code block
                </v-btn>
                <v-btn outlined small @click="editor.chain().focus().toggleBlockquote().run()"
                  :class="{ 'is-active': editor.isActive('blockquote') }">
                  blockquote
                </v-btn>
                <v-btn outlined small @click="editor.chain().focus().setHorizontalRule().run()">
                  horizontal rule
                </v-btn>
                <v-btn outlined small @click="editor.chain().focus().setHardBreak().run()">
                  hard break
                </v-btn>
                <v-btn outlined small @click="editor.chain().focus().undo().run()"
                  :disabled="!editor.can().chain().focus().undo().run()">
                  undo
                </v-btn>
                <v-btn outlined small @click="editor.chain().focus().redo().run()"
                  :disabled="!editor.can().chain().focus().redo().run()">
                  redo
                </v-btn>
              </v-container>
              <!-- <v-divider></v-divider> -->
            </v-col>
          </v-row>
          <v-row>
            <v-col>
              <editor-content :editor="editor" />
            </v-col>
          </v-row>
        <!-- </v-container> -->
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
      content: `
      <h2>
          Hi there,
        </h2>
        <p>
          this is a <em>basic</em> example of <strong>tiptap</strong>. Sure, there are all kind of basic text styles you’d probably expect from a text editor. But wait until you see the lists:
        </p>
        <ul>
          <li>
            That’s a bullet list with one …
          </li>
          <li>
            … or two list items.
          </li>
        </ul>
        <p>
          Isn’t that great? And all of that is editable. But wait, there’s more. Let’s try a code block:
        </p>
        <pre><code class="language-css">body {
  display: none;
}</code></pre>
        <p>
          I know, I know, this is impressive. It’s only the tip of the iceberg though. Give it a try and click a little bit around. Don’t forget to check the other examples too.
        </p>
        <blockquote>
          Wow, that’s amazing. Good work, boy! 👏
          <br />
          — Mom
        </blockquote>
      `,
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

<style lang="scss">
/* Basic editor styles */
.ProseMirror {
  > * + * {
    margin-top: 0.75em;
  }
  &:focus {
    outline: none;
  }

  min-height: 400px;

  ul,
  ol {
    padding: 0 1rem;
  }

  h1,
  h2,
  h3,
  h4,
  h5,
  h6 {
    line-height: 1.1;
  }

  code {
    background-color: rgba(#616161, 0.1);
    color: #616161;
  }

  pre {
    background: #0D0D0D;
    color: #FFF;
    font-family: 'JetBrainsMono', monospace;
    padding: 0.75rem 1rem;
    border-radius: 0.5rem;

    code {
      color: inherit;
      padding: 0;
      background: none;
      font-size: 0.8rem;
    }
  }

  img {
    max-width: 100%;
    height: auto;
  }

  blockquote {
    padding-left: 1rem;
    border-left: 2px solid rgba(#0D0D0D, 0.1);
  }

  hr {
    border: none;
    border-top: 2px solid rgba(#0D0D0D, 0.1);
    margin: 2rem 0;
  }
}
</style>