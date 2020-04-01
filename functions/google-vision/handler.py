# Copyright 2015 Google, Inc. 2019 TriggerMesh, Inc.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

import boto3
import json
import os

from google.cloud import vision
from google.cloud.vision import types
from PIL import Image, ImageDraw

s3 = boto3.resource('s3')
bucketname = os.environ['BUCKETNAME']
input_filename = '/tmp/foo.jpg'
output_filename = '/tmp/result.jpg'

max_results=4

def vision(event, context):
    print(json.dumps(event))
    print(type(event))
    print(event['body'])
    
    # [START vision_face_detection_tutorial_send_request]
    def detect_face(face_file, max_results=4):
        """Uses the Vision API to detect faces in the given file.
        Args:
            face_file: A file-like object containing an image with faces.
        Returns:
            An array of Face objects with information about the picture.
        """
        # [START vision_face_detection_tutorial_client]
        client = vision.ImageAnnotatorClient('/opt/key.json')
        # [END vision_face_detection_tutorial_client]

        content = face_file.read()
        image = types.Image(content=content)

        return client.face_detection(
            image=image, max_results=max_results).face_annotations
    # [END vision_face_detection_tutorial_send_request]
    # [START vision_face_detection_tutorial_process_response]
    def highlight_faces(image, faces, output_filename):
        """Draws a polygon around the faces, then saves to output_filename.
        Args:
          image: a file containing the image with the faces.
          faces: a list of faces found in the file. This should be in the format
              returned by the Vision API.
          output_filename: the name of the image file to be created, where the
              faces have polygons drawn around them.
        """
        im = Image.open(image)
        draw = ImageDraw.Draw(im)
        # Sepecify the font-family and the font-size
        for face in faces:
            box = [(vertex.x, vertex.y)
                   for vertex in face.bounding_poly.vertices]
            draw.line(box + [box[0]], width=5, fill='#00ff00')
            # Place the confidence value/score of the detected faces above the
            # detection box in the output image
            draw.text(((face.bounding_poly.vertices)[0].x,
                       (face.bounding_poly.vertices)[0].y - 30),
                      str(format(face.detection_confidence, '.3f')) + '%',
                      fill='#FF0000')
        im.save(output_filename)
    # [END vision_face_detection_tutorial_process_response]
    
    records=json.loads(event['body'])
    print(type(records))
    for rec in records['Records']:
        bucketname=rec['s3']['bucket']['name']
        name=rec['s3']['object']['key']
        print(bucketname)
        print(name)
        s3.Bucket(bucketname).download_file(name, input_filename)
        
    with open(input_filename, 'rb') as image:
        faces = detect_face(image, max_results)
        print('Found {} face{}'.format(
            len(faces), '' if len(faces) == 1 else 's'))

        print('Writing to file {}'.format(output_filename))
        # Reset the file pointer, so we can read the file again
        image.seek(0)
        highlight_faces(image, faces, output_filename)
    s3.Bucket(bucketname).upload_file(output_filename,'out.jpg')
    
    return
