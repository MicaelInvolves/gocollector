<template>
  <div id="app">
    <gc-navbar title="Go Collector"></gc-navbar>
    <gc-header title="Contatos registrados"></gc-header>

    <gc-container>
      <gc-contacts></gc-contacts>
    </gc-container>
  </div>
</template>

<script>
  import GcNavbar from './components/GcNavbar.vue'
  import GcHeader from './components/GcHeader.vue'
  import GcContainer from './components/GcContainer.vue'
  import GcContacts from './components/GcContacts.vue'

  export default {
    name: 'app',
    components: {
      GcNavbar,
      GcHeader,
      GcContainer,
      GcContacts
    },
    data() {
      return {
        msg: 'Bulma'
      }
    }
  }

</script>

<style>


</style>
