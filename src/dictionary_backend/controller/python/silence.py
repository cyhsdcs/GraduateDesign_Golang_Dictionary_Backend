from ffmpy3 import FFmpeg
import os


def silenceremove(filePath, fileName):

    inputFile = filePath+fileName
    outputFile = filePath+"temp_"+fileName

    try:
        ff = FFmpeg(inputs = {inputFile:None},
                    outputs = {outputFile:' -af silenceremove=stop_periods=-1:stop_duration=0.2:stop_threshold=-30dB'})
        print(ff.cmd)
        ff.run()
    except Exception as e:
        print(e)
        print('silence remove failed!\r\n')
    else:
        print('silence remove suceess!\r\n')


    try:
        os.remove(inputFile)
        os.rename(outputFile,inputFile)
    except Exception as e:
        print(e)
        print('remove and rename file failed!\r\n')
    else:
        print('rename file success\r\n')


