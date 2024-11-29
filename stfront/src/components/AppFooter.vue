<template>
  <v-footer
    height="40"
    app
  >
    <div
      class="text-caption text-disabled"
      style="position: absolute; right: 16px;"
    >
      &copy; {{ (new Date()).getFullYear() }} <span class="d-none d-sm-inline-block">Andey Robins</span>
      —
      <a
        class="text-decoration-none on-surface"
        href="https://github.com/andey-robins/space-traders/blob/main/LICENSE"
        rel="noopener noreferrer"
        target="_blank"
      >
        GPL v3 
      </a>
    </div>
  </v-footer>
</template>

<script setup lang="ts">

</script>

<style scoped lang="sass">
  .social-link :deep(.v-icon)
    color: rgba(var(--v-theme-on-background), var(--v-disabled-opacity))
    text-decoration: none
    transition: .2s ease-in-out

    &:hover
      color: rgba(25, 118, 210, 1)
</style>
