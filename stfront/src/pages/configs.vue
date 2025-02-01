<template>
  <v-container>
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
      <template #[`item.actionDelete`]="{ item }">
        <v-btn
          @click.prevent="del(item.id)"
        >
          Delete
        </v-btn>
      </template>
    </v-data-table>
    <v-btn
      @click.prevent="createNew"
    >
      New
    </v-btn>
  </v-container>
</template>

<script lang="ts">
  import { configService } from '../services/config-service'
  import { type Config, type ConfigId } from '../types/config'
  import { type DataTableHeader } from '../types/base'

  export default defineComponent({
    data: () => ({
      isLoading: false as boolean,

      configs: [] as Config[],
      headers: [
        { title: "Key", key: "key", width: 200 },
        { title: "Value", key: "value", width: 200 },
        { title: "Save", key: "actionSave", width: 25 },
        { title: "Delete", key: "actionDelete", width: 25 },
      ] satisfies DataTableHeader[],

      errorMessage: "" as string,
    }),
    mounted() {
      this.loadData();
    },
      methods: {
        catchError(err: Error): void {
          console.log(err)
          this.errorMessage = err.message;
        },
        createNew(): void {
          this.configs.push({id: -1, key: "", value: ""})
        },
        async loadData(): Promise<void> {
          this.isLoading = true;
          configService.getAll()
            .then((res: Config[]) => {
              this.configs = res
            }).catch(this.catchError)
            .finally(() => {
              this.isLoading = false;
            });
        },
        async save(dto: Config) {
          configService.save(dto)
            .then(() => {
              this.loadData();
            }).catch(this.catchError);
        },
        async del(id: ConfigId) {
          configService.del(id)
            .then(() => {
              this.loadData();
            }).catch(this.catchError);
        }
      },
    })
</script>