<script lang="ts">
	import '../app.pcss';
	import '$lib/app.css';
	import TitleBar from '$lib/components/TitleBar.svelte';
	import {onDestroy} from "svelte";
	import {currentSessionName, encryptionKey, isOffline, sessionNames, wasAutoLoggedOut} from "$lib/states";

	let autoLogoutOnBlur: number | undefined
	onDestroy(() => {
		clearTimeout(autoLogoutOnBlur)
	})

	let online: boolean
	$: {
		$isOffline = !online
		if ($isOffline) {
			logout()
		}
	}

	function logout() {
		if ($encryptionKey == null) return
		$encryptionKey = null
		$sessionNames  = null
		$currentSessionName = null
		$wasAutoLoggedOut = true
	}

	function onFocus() {
		clearTimeout(autoLogoutOnBlur)

		autoLogoutOnBlur = undefined
	}

	function onBlur() {
		autoLogoutOnBlur = setTimeout(logout, (1_000 * 60 * 5))
	}
</script>

<svelte:window bind:online on:focus={onFocus} on:blur={onBlur}/>
<div class="max-h-screen min-h-screen w-full px-8">
	<TitleBar></TitleBar>
	<div class="h-full w-full">
		<slot />
	</div>
</div>
