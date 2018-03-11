<template>
  <nav class="navbar is-dark">
    <div class="container">
      <div class="navbar-brand">
        <span class="navbar-item">
          {{ title }}
        </span>
      </div>
    </div>
  </nav>
</template>

<script>
export default {
  name: "GcNavbar",
  props: {
    title :{
      required: true,
      type: String
    }
  }
}
</script>
