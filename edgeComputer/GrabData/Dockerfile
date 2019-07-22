#This will act as container1 grabbing data with bash scripts
FROM alpine:3.10
RUN apk add --no-cache bash
RUN apk --no-cache add wget

#create MEGAGRONROCKS directory
RUN mkdir /MEGATRONROCKS

#create run wget command directory
RUN mkdir /MEGATRONROCKS/runWGET

#copy wow.sh script to MEGATRONROCKS
COPY /wow.sh /MEGATRONROCKS/runWGET/wow.sh

ENTRYPOINT ["bash","/MEGATRONROCKS/runWGET/wow.sh"]