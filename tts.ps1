Add-Type –AssemblyName System.Speech

$synth = New-Object System.Speech.Synthesis.SpeechSynthesizer

$synth.Speak("${env:TTS_TEXT}")
