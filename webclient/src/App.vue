<template>
  <div id="app">
    <nav class="navbar is-dark">
      <div class="container">
        <div class="navbar-brand">
          <span class="navbar-item">
            Go Collector
          </span>
        </div>
      </div>
    </nav>
    <div class="hero is-primary is-bold">
      <div class="hero-body">
        <div class="container">
          <h1 class="title">Contatos</h1>
        </div>
      </div>
    </div>
    <section class="section">
      <div class="container">
        <div class="buttons is-right">
          <span class="button is-primary">Atualizar</span>
        </div>
      </div>
      <div class="container">
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
            <tr>
              <td>Jay</td>
              <td>exemplo@email.com</td>
              <td>50</td>
              <td>home, contact, prices</td>
            </tr>
          </tbody>
        </table>
      </div>
    </section>
  </div>
</template>

<script>
  export default {
    name: 'app',
    data() {
      return {
        msg: 'Bulma'
      }
    }
  }

</script>

<style>


</style>
