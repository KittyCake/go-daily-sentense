<template>
	<div id="app" class="container">
		<div class="row">
			<div class="col-md-6 offset-md-3 py-5">
				<div>
					<button class="btn btn-primary" @click="generate">Generate!</button>
				</div>
				<h2>{{ dailySentense }}</h2>
				<img :src="dailyImg" />
			</div>
		</div>
	</div>
</template>

<script>
	import axios from 'axios'
	import _ from 'lodash'

	export default {
		name: 'App',
		data() {
			return {
				dailySentense: '',
				dailyImg: '',
				imgArray: [
					'https://lh3.googleusercontent.com/proxy/fQ11ftSykkl-mAK0hs3oNiCjVt3Ms3ujCZPVzNVslsumpkorlGbSxe7v8wLLKelMeRMzE2Qmtpy6NlGx2t3sJRbrO5m0aP4clZxTvq-HDqUwdYJPxsiWbQ',
					'https://inews.gtimg.com/newsapp_match/0/9716133960/0',
					'https://lh3.googleusercontent.com/proxy/f8-xbz1Ek8XPwonMc2yprC2MrqsP3R-1VUOIwbnDTJXVb7uWEqXpWoara43n_nO-yyt5LzUoi5YuwGKFePxHp_ewfJgG5zhgFNsfe9kiozGnts0f9E8MgA',
					'https://cms-bucket.nosdn.127.net/2018/10/29/08b54f3ebf89433b9ffea9e0f1900211.jpeg',
					'https://5b0988e595225.cdn.sohucs.com/images/20190701/642e35274c0f4892a309e65e4353fe8f.gif',
					'https://encrypted-tbn0.gstatic.com/images?q=tbn%3AANd9GcTk3OVB3RdBwbrKBSkVZln9040422gJFM57OA&usqp=CAU',
				],
			}
		},
		components: {},
		methods: {
			async generate() {
				await axios
					.get('http://localhost:3000/api/sentence')
					.then((response) => {
						this.dailySentense = response.data.slip.advice
					})

				this.dailyImg = _.shuffle(this.imgArray)[0]
			},
		},
	}
</script>

<style>
	#app {
		font-family: Avenir, Helvetica, Arial, sans-serif;
		-webkit-font-smoothing: antialiased;
		-moz-osx-font-smoothing: grayscale;
		text-align: center;
		color: #2c3e50;
		margin-top: 60px;
	}
	button,
	h2,
	img {
		margin-bottom: 20px;
	}
</style>
