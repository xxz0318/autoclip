// Package audio 请修改此处包名注释
// @author: xiexinzhong
// @create: 2024-01-26 22:01
// @description:
package audio

import (
	"context"
	"reflect"
	"testing"

	"douyin_video/conf"
	"douyin_video/log"
)

func init() {
	conf.LoadConfig()
	log.InitLog()
}
func TestTxtToAudio(t *testing.T) {
	type args struct {
		ctx           context.Context
		content       string
		audioFileName string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "",
			args: args{
				ctx:           context.Background(),
				content:       "穿越女当众作诗：“碧玉妆成一树高，万条垂下绿丝绦……”\n我默念：「不知细叶谁裁出，二月春风似剪刀。这是贺知章的诗吧，几年级学的，穿书太多年，忘了。」\n谢琅忽然看向我。\n我不动声色腹诽：「看我干什么，看你的女主角啊！」\n我穿进一本书里，男主是谢琅，女主是我的表姐——一个穿越女。\n按照剧情，谢琅会和表姐HE。\n可书里没说，他能读我的心啊。\n1\n侯府世子在生辰宴上提议：“咱们今日，就以酒为题每人做一首诗如何？”\n我心知肚明，表姐小有才名，最擅诗词——这是他给未婚妻露脸的机会。\n可惜，他只是男二。\n表姐就坐在我身旁，提笔前对我说：“妹妹若是不会，还是不要勉强。”\n得，我还是扮演好胸无点墨的愚蠢表妹吧。\n正想着，谢琅到了。\n身为男主，那身皮囊自然是极好。\n他的目光朝着我略有停顿，随后落座。\n表姐痴愣愣望着他，眼里闪过隐晦的惊艳之色。\n原来是这俩苦命鸳鸯在对望，我就说……\n霎时间，我觉得头上绿油油的。\n因为……男主现在跟我定有婚约。\n说来也是无奈，商户谢家因着祖上的救命之恩，与表姐订有婚约，但姑母心高气傲，强行退了婚。\n而我和谢琅的婚事，则是姑母为了打发谢琅安排的。\n2\n一炷香时间到，表姐隐晦地往谢琅的方向扫一眼，“世子，妹妹那首我替她作了。”\n我默默鼓掌，不愧是女主，抓住一切机会出风头。\n众人都朝我和表姐看来。\n“竟然还可以替写，这妹妹原来是个草包。”\n“我一首都要抠掉头发了，林婉儿还能做两首？”\n“数量不重要，质量才是最重要的。”\n侍郎之女刘缨喜欢侯府世子，讥讽道：“我倒要看看她做了两首什么样的，可别是托大。”\n就是这些嘲笑林婉儿的，等会就要被啪啪打脸，并对她佩服得五体投地。\n我顺理成章交了白卷，默默看戏。\n至于骂我草包……\n我又不是古人，不会吟诗作对很正常。\n而且，我宁可被骂草包，也不想成为“文抄公”，把别人呕心沥血的作品信手拈来，装点自己的才女名声。\n侯府世子在一堆诗词中翻看，忽然眼睛一亮，念出声来：“‘花间一壶酒，独酌无相亲。举杯邀明月，对影成三人。’好诗！这是谁写的？”\n表姐施施然起身，“小女不才，这首乃替妹妹所作。”\n众人品味后道：\n“如此好诗，竟然是替妹妹所作？”\n“那自己的那一首，可不得好上天了。”\n当然，也有不同的声音。\n刘缨挑刺：“我怎么觉得这首诗不够完整……”\n侯府世子则帮腔：“婉儿在如此短时间内连作两首，属实不易，哪怕是半首，也足够惊艳在场众人。”\n场中众人点头附和。\n刘缨咬咬唇，再不甘也不得不闭嘴。\n我剥了个橘子安静吃着，心想：「本来就是半首，剩下十言想来表姐是背不全。旁人就算觉得这首诗不完整，也不会说出来扫世子雅兴，这刘缨还是太心急了，难怪会被打脸。」\n察觉到一道打量的目光，我顺着感觉探过去。\n可谢琅正目视前方，",
				audioFileName: "./test1.wav",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := TxtToAudio(tt.args.ctx, tt.args.content, tt.args.audioFileName); (err != nil) != tt.wantErr {
				t.Errorf("TxtToAudio() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAesDecrypt(t *testing.T) {
	type args struct {
		crypted string
		key     string
		iv      string
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				crypted: "g2Ts2f59KhPf6TpuqI6SbuH7qJAmbj0PfmspkrDvCc2msI9aKnALGiUhCyTIL9Yd7xW5UPWW1hH6hgRDfXdTO3Pag3FiGofVVDBhCQbdgnuXMAXal1rQ++8GWB0Uhrj86aZjR1iWyP7ODGXe6JbUyQ==",
				key:     "abcdefgabcdefg12",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := aesDecrypt(tt.args.crypted, tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("AesDecrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AesDecrypt() got = %v, want %v", got, tt.want)
			}
		})
	}
}
