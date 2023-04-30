package sox

import (
	"math"
	"os"
	"strings"
	"testing"

	log "github.com/schollz/logger"
	"github.com/stretchr/testify/assert"
)

func TestPiko(t *testing.T) {
	log.SetLevel("trace")
	fname := "amen_beats8_bpm172.wav"
	beats, bpm, err := GetBPM(fname)
	assert.Nil(t, err)
	assert.Equal(t, 8.0, beats)
	assert.Equal(t, 165.0, bpm)
}

func TestReverseReverb(t *testing.T) {
	log.SetLevel("trace")
	fname := "amen_beats8_bpm172.wav"
	fname2, err := ReverseReverb(fname, 7, 3)
	assert.Nil(t, err)
	if fname2 != fname {
		os.Rename(fname2, "test.wav")
	}
	Clean()
}
func TestRun(t *testing.T) {
	log.SetLevel("trace")
	stdout, stderr, err := run("sox", "--help")
	assert.Nil(t, err)
	assert.True(t, strings.Contains(stdout, "SoX"))
	assert.Empty(t, stderr)
}

func TestLength(t *testing.T) {
	length, err := Length("amen_beats8_bpm172.wav")
	assert.Nil(t, err)
	assert.Equal(t, 1.133354, length)
}

func TestInfo(t *testing.T) {
	samplerate, channnels, err := Info("amen_beats8_bpm172.wav")
	assert.Nil(t, err)
	assert.Equal(t, 48000, samplerate)
	assert.Equal(t, 2, channnels)
}

func TestSilence(t *testing.T) {
	fname2, err := SilenceAppend("amen_beats8_bpm172.wav", 1)
	assert.Nil(t, err)
	length1, _ := Length("amen_beats8_bpm172.wav")
	length2, _ := Length(fname2)
	assert.Less(t, math.Abs(length2-length1-1), 0.00001)

	fname2, err = SilencePrepend("amen_beats8_bpm172.wav", 1)
	assert.Nil(t, err)
	length1, _ = Length("amen_beats8_bpm172.wav")
	length2, _ = Length(fname2)
	assert.Less(t, math.Abs(length2-length1-1), 0.00001)

	fname3 := MustString(SilenceTrim(fname2))
	length3 := MustFloat(Length(fname3))
	assert.Greater(t, length2-length3, 1.0)

	os.Rename(fname3, "test.wav")
}

func TestTrim(t *testing.T) {
	var fname2 string
	var err error
	fname2, err = Trim("amen_beats8_bpm172.wav", 0.5, 0.5)
	assert.Nil(t, err)
	assert.Equal(t, 0.5, MustFloat(Length(fname2)))
	fname2, err = Trim("amen_beats8_bpm172.wav", 0.5)
	assert.Nil(t, err)
	assert.Equal(t, 0.633354, MustFloat(Length(fname2)))
}

func TestReverse(t *testing.T) {
	var fname2 string
	var err error
	fname2, err = Reverse("amen_beats8_bpm172.wav")
	assert.Nil(t, err)
	assert.Equal(t, MustFloat(Length("amen_beats8_bpm172.wav")), MustFloat(Length(fname2)))
}

func TestPitch(t *testing.T) {
	var fname2 string
	var err error
	fname2, err = Pitch("amen_beats8_bpm172.wav", 3)
	assert.Nil(t, err)
	assert.Equal(t, MustFloat(Length("amen_beats8_bpm172.wav")), MustFloat(Length(fname2)))
	os.Rename(fname2, "test.wav")
}

func TestJoin(t *testing.T) {
	var fname2 string
	var err error
	fname2, err = Join("amen_beats8_bpm172.wav", "amen_beats8_bpm172.wav", "amen_beats8_bpm172.wav")
	assert.Nil(t, err)
	assert.LessOrEqual(t, math.Abs(MustFloat(Length(fname2))-3*MustFloat(Length("amen_beats8_bpm172.wav"))), 0.001)
	os.Rename(fname2, "test.wav")
}

func TestRepeat(t *testing.T) {
	var fname2 string
	var err error
	fname2, err = Repeat("amen_beats8_bpm172.wav", 2)
	assert.Nil(t, err)
	assert.LessOrEqual(t, math.Abs(MustFloat(Length(fname2))-3*MustFloat(Length("amen_beats8_bpm172.wav"))), 0.001)
	os.Rename(fname2, "test.wav")
}

func TestRetempoSpeed(t *testing.T) {
	var fname2 string
	var err error
	fname2, err = RetempoSpeed("amen_beats8_bpm172.wav", 60, 120)
	assert.Nil(t, err)
	assert.LessOrEqual(t, math.Abs(MustFloat(Length("amen_beats8_bpm172.wav"))/2-MustFloat(Length(fname2))), 0.001)
}
func TestRetempoStretch(t *testing.T) {
	var fname2 string
	var err error
	fname2, err = RetempoStretch("amen_beats8_bpm172.wav", 60, 120)
	assert.Nil(t, err)
	assert.LessOrEqual(t, math.Abs(MustFloat(Length("amen_beats8_bpm172.wav"))/2-MustFloat(Length(fname2))), 0.001)
}

func TestCopyPaste(t *testing.T) {
	var fname2 string
	var err error
	fname2, err = CopyPaste("amen_beats8_bpm172.wav", 0.14, 0.27, 0.57, 0.02)
	assert.Nil(t, err)
	assert.Equal(t, MustFloat(Length("amen_beats8_bpm172.wav")), MustFloat(Length(fname2)))
	os.Rename(fname2, "test.wav")
}

func TestPaste(t *testing.T) {
	var fname2 string
	var err error
	crossfade := 0.04
	piece := MustString(Trim("amen_beats8_bpm172.wav", 0.14-crossfade, 0.27+crossfade))
	fname2, err = Paste("amen_beats8_bpm172.wav", piece, 0.57, crossfade)
	assert.Nil(t, err)
	assert.Equal(t, MustFloat(Length("amen_beats8_bpm172.wav")), MustFloat(Length(fname2)))
	os.Rename(fname2, "test.wav")
}

func TestGain(t *testing.T) {
	var fname2 string
	var err error
	fname2, err = Gain("amen_beats8_bpm172.wav", 6)
	assert.Nil(t, err)
	assert.Equal(t, MustFloat(Length("amen_beats8_bpm172.wav")), MustFloat(Length(fname2)))
	os.Rename(fname2, "test.wav")
}

func TestSampleRate(t *testing.T) {
	var fname2 string
	var err error
	fname2, err = SampleRate("amen_beats8_bpm172.wav", 8000)
	assert.Nil(t, err)
	assert.Less(t, math.Floor(MustFloat(Length("amen_beats8_bpm172.wav"))-MustFloat(Length(fname2))), 0.001)
	os.Rename(fname2, "test.wav")
}

func TestStretch(t *testing.T) {
	var fname2 string
	var err error
	fname := "amen_beats8_bpm172.wav"

	fname2, err = Stretch(fname, 2)
	assert.Nil(t, err)
	assert.Less(t, math.Abs(MustFloat(Length(fname))*2-MustFloat(Length(fname2))), 0.01)
	os.Rename(fname2, "test.wav")
}

func TestStutter(t *testing.T) {
	var fname2 string
	var err error
	fname2, err = Stutter("amen_beats8_bpm172.wav", 60.0/160/4, 0.5, 4, 0.005)
	assert.Nil(t, err)
	if fname2 != "amen_beats8_bpm172.wav" {
		os.Rename(fname2, "test.wav")
	}
}

// keep this last
func TestClean(t *testing.T) {
	assert.Nil(t, Clean())
}