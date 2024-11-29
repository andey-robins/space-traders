<template>
  <v-data-table
    :items="configs"
    :headers="headers"
    no-data-text="No configs loaded"
  >
    <template #[`item.key`]="{ item }">
      <v-text-field
        v-model="item.key"
      />
    </template>
    <template #[`item.value`]="{ item }">
      <v-text-field
        v-model="item.value"
      />
    </template>
    <template #[`item.actionSave`]="{ item }">
      <v-btn
        @click.prevent="save(item)"
      >
        Save
      </v-btn>
    </template>
  </v-data-table>
</template>

<script lang="ts">
    import { configService } from '@/services/config-service.ts'

    export default defineComponent({
        data: () => ({
            isLoading: false as boolean,

            configs: [] as Config[],
            headers: [
                { title: "Key", key: "key" },
                { title: "Value", value: "value" },
                { title: "Save", value: "actionSave" },
            ] as DataTableHeaders[],

            errorMessage: "",
        }),
        mounted() {
            this.loadData();
        },
        methods: {
            catchError(err: error): void {
                this.errorMessage = err;
            },
            async loadData(): void {
                this.isLoading = true;
                configService.getAll()
                    .then((res: Config[]) => {
                        this.configs = res.data
                    }).catch(this.catchError)
                    .finally(() => {
                        this.isLoading = false;
                    });
            },
            async save(dto: Config) {
                configService.save(dto)
                  .then(() => {
                    this.loadData();
                  });
            },
        },
    })
</script>