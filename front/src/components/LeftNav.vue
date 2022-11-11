<template>

  <v-navigation-drawer
      v-model="navigationDrawer"
      app
      
      
    >
      <v-sheet
        class="pa-4"
        
      >
      <v-row>
        <v-col>

          <v-avatar
          class="mb-4"
          size="64"
          color="blue lighten-2"
          >LS</v-avatar>
        </v-col>
        <v-col>
          <v-switch v-model="$vuetify.theme.dark"></v-switch>
        </v-col>
      </v-row>

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
      <v-spacer></v-spacer>
    </v-navigation-drawer>
</template>

<script>
import settings from '@/settings';
import vuetify from '@/plugins/vuetify';

export default {
    name: "LeftNav",
    props: ['navigationDrawer'],
    data: () =>({
      bucketList: [],
      bucketName: '',
      drawer: true,
    }),
    methods: {
      GetBucketList() {
        console.log(location.href)
        this.$http.get(settings.ebaHOST + "api/bucket/list").then((resp) => {
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
