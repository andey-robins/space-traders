<template>
  <div>
    <v-app-bar 
      :elevation="2"
      prominent
    >
      <template #prepend>
        <v-app-bar-nav-icon @click.stop="drawer = !drawer" />
      
        <v-app-bar-title>
          Space Traders
        </v-app-bar-title>
      </template>
    </v-app-bar>
    <v-navigation-drawer
      v-model="drawer"
      temporary
    >
      <v-list-item title="System" />
      <v-list-item
        href="/"
        title="Dashboard"
      />
      <v-list-item 
        href="/configs" 
        title="Configs" 
      />
    </v-navigation-drawer>
  </div>
</template>

<script lang="ts">
  export default defineComponent({
    data: () => ({
      drawer: false
    })
  })
</script>
