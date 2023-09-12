
<script>
	import { fade } from "svelte/transition";
</script>
<div class="modal" transition:fade={{ delay: 0, duration: 150 }}>
	<slot></slot>
</div>
<style>
.modal {
	background-color: rgba(0, 0, 0, 0.8);
	width: 100%;
	position: absolute;
	z-index: 10;
	height: calc(100% - 36px);
}
</style>