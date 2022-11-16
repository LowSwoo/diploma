<template>
    <v-dialog transition="dialog-top-transition" max-width="600px" v-model="showDescription">
        <v-card>
            <v-card-title> Здесь должно быть описание расчёта</v-card-title>
            <v-card-text> Ну и тут штука от типтапа</v-card-text>
            <v-card-actions>
                <v-btn @click="$emit('CloseDescriptionDialog')"/>
            </v-card-actions>
        </v-card>
    </v-dialog>
</template>


<script>

export default {
    name: "showDescriptionDialog",
    props: ['showDescription'],
}

</script>