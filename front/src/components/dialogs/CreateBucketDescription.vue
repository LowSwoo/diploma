<template>
  <v-dialog v-model="showCreateDescriptionDialog" fullscreen hide-overlay transition="dialog-bottom-transition"
    scrollable>
    
    <v-card tile>
      
      <v-toolbar flat dark color="primary">
        <v-toolbar-title height="200">Описание расчёта</v-toolbar-title>
        <v-btn icon dark @click="CloseDescriptionDialog">
          <v-icon>mdi-close</v-icon>
        </v-btn>
        <v-spacer></v-spacer>
        <v-toolbar-items>
          <v-btn dark text @click="dialog = false">
            Save
          </v-btn>
        </v-toolbar-items>
      </v-toolbar>
      <v-card-text>
        <v-sheet min-height="90vh" rounded="lg" elevation="4">
          <div id="editor">
        <v-row >
          <v-col cols="6">
            <textarea autoGrow background-color="#424242" :value="input" @input="update"></textarea>
          </v-col>
          <v-col cols="6">
            <div v-html="markDownToHtml"></div>
          </v-col>
        </v-row>
      </div>
      </v-sheet>       
      </v-card-text>

      <div style="flex: 1 1 auto;"></div>
    </v-card>
  </v-dialog>
</template>

<script>
import { marked } from "marked"
import lodash from "lodash"
export default {
  name: "CreateBucketDescription",
  props: ["showCreateDescriptionDialog"],
  data: () => ({
    input: "# Заголовок первого уровня"
  }),
  computed: {
    markDownToHtml() {
      return marked(this.input)
    }
  },
  methods: {
    CloseDescriptionDialog() {
      this.$emit('CloseDescriptionDialog')
    },
    update: _.debounce(function (e) {
      this.input = e.target.value
    }, 300)
  }
}

</script>
<style>
html,
body,
#editor {
  margin: 0;
  height: 100%;
  font-family: "Helvetica Neue", Arial, sans-serif;
  color: #333;
}

textarea,
#editor div {
  display: inline-block;
  width: 49%;
  height: 100%;
  vertical-align: top;
  box-sizing: border-box;
  padding: 0 20px;
}

textarea {
  border: none;
  border-right: 1px solid #ccc;
  resize: none;
  outline: none;
  background-color: #f6f6f6;
  font-size: 14px;
  font-family: "Monaco", courier, monospace;
  padding: 20px;
}

code {
  color: #f66;
}

</style>