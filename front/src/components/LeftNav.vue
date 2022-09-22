<template>
  <!-- <v-col cols="2">
    <v-sheet rounded="lg" elevation="2">
      <v-list color="transparent">
        <v-list-item>
            <v-list-item>Текущие расчёты</v-list-item>
        </v-list-item>
        <v-divider class="my-2"></v-divider>

        <v-list-item v-for="bucket in bucketList" :key="bucket" link>
          <v-list-item-content v-on:click="$root.$emit('ChangeCurrentBucket', bucket)">
            <v-list-item-title> {{ bucket }} </v-list-item-title>
          </v-list-item-content>
        </v-list-item>
      </v-list>
    </v-sheet>
  </v-col> -->

  <v-navigation-drawer
      v-model="drawer"
      app
      color="blue-grey darken-4"
    >
      <v-sheet
        class="pa-4"
        color="blue-grey darken-4"
      >
        <v-avatar
          class="mb-4"
          size="64"
          color="blue lighten-2"
        >LS</v-avatar>

        <!-- <div>github.com/lowswoo</div> -->
      </v-sheet>

      <v-divider></v-divider>

      <v-list >
        <v-list-item v-if="bucketList.length != 0">Доступные расчёты</v-list-item>
        <v-list-item v-else>Здесь пока нет расчётов</v-list-item>
        <v-divider class="my-2"></v-divider>
        <v-list-item
          v-for="bucket in bucketList"
          link
        >
          <v-list-item-icon>
            <v-icon>mdi-calculator</v-icon>
          </v-list-item-icon>

          <v-list-item-content v-on:click="$root.$emit('ChangeCurrentBucket', bucket)">
            <v-list-item-title>{{ bucket }}</v-list-item-title>
          </v-list-item-content>
        </v-list-item>
      </v-list>
    </v-navigation-drawer>
</template>

<script>

export default {
    name: "LeftNav",
    data: () =>({
      bucketList: [],
      bucketName: '',
      drawer: null,
    }),
    methods: {
      GetBucketList() {
        this.$http.get("http://localhost:8080/api/bucket/list").then((resp) => {
        this.bucketList = resp.data;
        this.bucketName = this.bucketList[0]
        this.$root.$emit('ChangeCurrentBucket', this.bucketName)
      })
      }
    },
    created() {
      this.GetBucketList()
    },
    mounted() {
      this.$root.$on('UpdateBucketList', () => this.GetBucketList())
    },

}


</script>
