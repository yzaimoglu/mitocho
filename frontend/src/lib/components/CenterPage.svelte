<script lang="ts">
  export let gap = 4;
</script>

<div class={`flex flex-col justify-center items-center min-h-screen gap-${gap}`}>
  <slot />
</div>
