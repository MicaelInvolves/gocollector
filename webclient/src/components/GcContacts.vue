<template>
  <div id="contacts">
    <div class="buttons is-right">
      <span class="button is-primary" @click="updateContacts()">Atualizar</span>
    </div>
    <table class="table is-striped is-fullwidth">
      <thead>
        <tr>
          <th>Nome</th>
          <th>Email</th>
          <th>Visitas</th>
          <th>Paginas</th>
        </tr>
      </thead>
      <tfoot>
        <tr>
          <th>Nome</th>
          <th>Email</th>
          <th>Visitas</th>
          <th>Paginas</th>
        </tr>
      </tfoot>
      <tbody>
        <tr v-for="contact in contacts" :key="contact.id">
          <td>{{ contact.name }}</td>
          <td>{{ contact.email }}</td>
          <td>{{ contact.visits }}</td>
          <td>{{ contact.pages }}</td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script>
export default {
  name: 'GcContacts',
  data() {
    return {
      contacts: []
    }
  },
  created() {
    this.updateContacts();
  },
  methods: {
    updateContacts() {
      this.$http.get('/api/subscribers').then(result => {
        this.contacts = result.data;
      }).catch(result => {
        console.log(result);
        this.contacts = [];
      });
    }
  }
}
</script>
